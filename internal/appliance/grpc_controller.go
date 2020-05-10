package appliance

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/gregorychen3/espresso-controller/internal/appliance/heating_element"
	"github.com/gregorychen3/espresso-controller/internal/appliance/temperature"
	"github.com/gregorychen3/espresso-controller/pkg/appliancepb"
	"github.com/gregorychen3/espresso-controller/pkg/control"
	"github.com/gregorychen3/espresso-controller/pkg/control/bangbang"
	"github.com/pkg/errors"
)

type TemperatureController interface {
	control.Strategy
	Shutdown() error
}

type grpcController struct {
	c                      Configuration
	boilerTemperatureCtrlr TemperatureController
	groupTherm             temperature.Sampler
}

func newGrpcController(
	c Configuration,
	heatingElem heating_element.HeatingElement,
	boilerTherm temperature.Sampler,
	groupTherm temperature.Sampler,
) (*grpcController, error) {
	temperatureCtrlr, err := bangbang.NewBangbang(heatingElem, boilerTherm)
	if err != nil {
		return nil, err
	}

	if err := temperatureCtrlr.Run(); err != nil {
		return nil, errors.Wrap(err, "Failed to start temperature controller")
	}

	return &grpcController{
		c:                      c,
		boilerTemperatureCtrlr: temperatureCtrlr,
		groupTherm:             groupTherm,
	}, nil
}

func (c *grpcController) BoilerTemperature(req *appliancepb.TemperatureStreamRequest, stream appliancepb.Appliance_BoilerTemperatureServer) error {
	// the first message sent on the stream is the temperature history
	var pbSamples []*appliancepb.TemperatureSample
	samples := c.boilerTemperatureCtrlr.GetTemperatureHistory()
	for _, s := range samples {
		pbTime, err := ptypes.TimestampProto(s.ObservedAt)
		if err != nil {
			return err
		}
		pbSample := appliancepb.TemperatureSample{
			Value:      s.Value,
			ObservedAt: pbTime,
		}
		pbSamples = append(pbSamples, &pbSample)
	}

	if err := stream.Send(&appliancepb.TemperatureStreamResponse{
		Data: &appliancepb.TemperatureStreamResponse_History{
			History: &appliancepb.TemperatureHistory{
				Samples: pbSamples,
			},
		},
	}); err != nil {
		return err
	}

	// send a current sample every second
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			sample := c.boilerTemperatureCtrlr.GetCurrentTemperature()
			pbTime, err := ptypes.TimestampProto(sample.ObservedAt)
			if err != nil {
				return err
			}
			if err := stream.Send(&appliancepb.TemperatureStreamResponse{
				Data: &appliancepb.TemperatureStreamResponse_Sample{
					Sample: &appliancepb.TemperatureSample{
						Value:      sample.Value,
						ObservedAt: pbTime,
					},
				},
			}); err != nil {
				return err
			}
		}
	}
}

func (c *grpcController) GetTargetTemperature(context.Context, *appliancepb.GetTargetTemperatureRequest) (*appliancepb.GetTargetTemperatureResponse, error) {
	targetTemperature := c.boilerTemperatureCtrlr.GetTargetTemperature()

	pbTime, err := ptypes.TimestampProto(targetTemperature.SetAt)
	if err != nil {
		return nil, err
	}

	return &appliancepb.GetTargetTemperatureResponse{
		Temperature: targetTemperature.Value,
		SetAt:       pbTime,
	}, nil
}

func (c *grpcController) SetTargetTemperature(ctx context.Context, req *appliancepb.SetTargetTemperatureRequest) (*appliancepb.SetTargetTemperatureResponse, error) {
	targetTemperature := c.boilerTemperatureCtrlr.SetTargetTemperature(req.Temperature)

	pbTime, err := ptypes.TimestampProto(targetTemperature.SetAt)
	if err != nil {
		return nil, err
	}

	return &appliancepb.SetTargetTemperatureResponse{
		Temperature: targetTemperature.Value,
		SetAt:       pbTime,
	}, nil
}

func (c *grpcController) Shutdown() error {
	return c.boilerTemperatureCtrlr.Shutdown()
}
