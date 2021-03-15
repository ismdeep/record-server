package serializer

import (
	"record-server/model"
	"time"
)

type Record struct {
	ID        uint      `json:"id"`
	EventTime time.Time `json:"event_time"`
	Amount    float64   `json:"amount"`
	Comment   string    `json:"comment"`
	Tags      string    `json:"tags"`
}

func BuildRecord(record model.Record) Record {
	return Record{
		ID:        record.ID,
		EventTime: record.EventTime,
		Amount:    record.Amount,
		Comment:   record.Comment,
		Tags:      record.Tags,
	}
}

func BuildRecordResponse(record model.Record) Response {
	return Response{
		Data: BuildRecord(record),
	}
}

func BuildRecordListResponse(records []model.Record) Response {
	var data []Record

	for _, record := range records {
		data = append(data, BuildRecord(record))
	}

	return Response{
		Data: data,
	}
}
