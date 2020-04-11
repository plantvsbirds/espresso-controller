package pid

import "time"

type SetPoint struct {
	Temperature float32
	SetAt       time.Time
}

type PID struct {
	setPoints []SetPoint
}

func NewPID() *PID {
	return &PID{
		setPoints: []SetPoint{{
			Temperature: 93,
			SetAt:       time.Now(),
		}},
	}
}

func (p *PID) GetCurrentTemperature() float32 {
	return 21.5
}

func (p *PID) GetSetPoint() SetPoint {
	return p.setPoints[len(p.setPoints)-1]
}

func (p *PID) SetSetPoint(temperature float32) SetPoint {
	setPoint := SetPoint{
		Temperature: temperature,
		SetAt:       time.Now(),
	}
	p.setPoints = append(p.setPoints, setPoint)
	return setPoint
}
