package bangbang

import (
	"math/rand"
	"sync"
	"time"

	"github.com/gregorychen3/espresso-controller/internal/log"
	"github.com/gregorychen3/espresso-controller/pkg/control"
	"go.uber.org/zap"
)

const (
	min float32 = 80.0
	max float32 = 100.0
)

// Bangbang is a temperature controller that implements bang-bang control.
// https://en.wikipedia.org/wiki/Bang%E2%80%93bang_control
type Bangbang struct {
	setPoints []control.SetPoint

	temperatureHistoryMu sync.RWMutex
	temperatureHistory   []control.TemperatureSample
}

func NewBangbang() *Bangbang {
	return &Bangbang{
		setPoints: []control.SetPoint{{
			Temperature: 93,
			SetAt:       time.Now(),
		}},
	}
}

func (p *Bangbang) Run() error {
	go func() {
		for {
			sample := p.sampleTemperature()
			p.temperatureHistoryMu.Lock()
			p.temperatureHistory = append(p.temperatureHistory, sample)
			p.temperatureHistoryMu.Unlock()
			time.Sleep(5 * time.Second)
		}
	}()

	return nil
}

func (p *Bangbang) GetCurrentTemperature() control.TemperatureSample {
	p.temperatureHistoryMu.RLock()
	defer p.temperatureHistoryMu.RUnlock()

	return p.temperatureHistory[len(p.temperatureHistory)-1]
}

func (p *Bangbang) GetTemperatureHistory() []control.TemperatureSample {
	p.temperatureHistoryMu.RLock()
	defer p.temperatureHistoryMu.RUnlock()

	return p.temperatureHistory
}

func (p *Bangbang) GetSetPoint() control.SetPoint {
	return p.setPoints[len(p.setPoints)-1]
}

func (p *Bangbang) SetSetPoint(temperature float32) control.SetPoint {
	setPoint := control.SetPoint{
		Temperature: temperature,
		SetAt:       time.Now(),
	}
	p.setPoints = append(p.setPoints, setPoint)
	return setPoint
}

func (p *Bangbang) sampleTemperature() control.TemperatureSample {
	randTemp := min + rand.Float32()*(max-min)
	log.Debug("Temperature sampled", zap.Float32("temperature", randTemp))
	return control.TemperatureSample{Value: randTemp, ObservedAt: time.Now()}
}
