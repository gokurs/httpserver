package sensor

// ErrReadSensor beschreibt den Fehler, wenn von einem Sensor ein lesen nicht
// möglich ist
var ErrReadSensor error

// nichtExportiert benötigt keine Dokumentation, jedoch man schreibt
// denk Code damit der gelesen und verstanden wird
func nichtExportiert() {

}

// Sensor is the interface for reading from a sensor
type Sensor interface {
	ID()
	Read(p []byte) (n int, err error)
	Close()
}

// Read verwendet einen Sensor und liest die Daten aus.
func Read(s Sensor) {}
