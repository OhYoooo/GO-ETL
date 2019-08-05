package pipelines

import (
	"common-etl/models"
	"common-etl/processors"
	"common-etl/readers"
	"common-etl/writers"
)

// Pipeline encapsulates the pipeline build for ETL process
type Pipeline struct {
	reader          *readers.Reader
	processor       *processors.Processor
	writer          *writers.Writer
	transInputChan  chan []models.PubsubModel
	transOutputChan chan []models.DatastoreEntity
}

// NewPipeline : get instance of build-up pipeline
func NewPipeline() *Pipeline {
	return &Pipeline{
		readers.NewReader(),
		processors.NewProcessor(),
		writers.NewWriter(),
		make(chan []models.PubsubModel, 2),
		make(chan []models.DatastoreEntity, 2),
	}
}

// Start the pipeline
func (p *Pipeline) Start() {
	go p.reader.Extract(p.transInputChan)
	go p.processor.ProcessData(p.transInputChan, p.transOutputChan)
	go p.writer.Write(p.transOutputChan)
}

// Stop the pipeline
// TODO: remove done <- true when more pipelines are created
func (p *Pipeline) Stop(done chan bool) {
	p.reader.Stop()
	done <- true
}
