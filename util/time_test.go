package util

import (
	"testing"
	"time"
)

func TestStringToTime(t *testing.T) {
	var convertdTime = StringToTime("2019-02-12T10:10:10")

	if convertdTime.Year() != 2019 {
		t.Errorf("Espera que o ano sejá 2019")
	}

	if convertdTime.Month() != time.February {
		t.Errorf("Espera que o mes sejá Fevereiro")
	}

	if convertdTime.Day() != 12 {
		t.Errorf("Esperado que o dia sejá 12")
	}

	if convertdTime.Hour() != 10 {
		t.Errorf("Esperado que a Hora sejá 10")
	}

	if convertdTime.Minute() != 10 {
		t.Errorf("Esperado que os Minutos sejá 10")
	}
}
