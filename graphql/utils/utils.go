package utils

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func NullablePgTypeText(Args map[string]interface{}, field string) pgtype.Text {
	if val, ok := Args[field]; ok && val != nil {
		return pgtype.Text{
			Valid:  Args[field].(string) != "",
			String: Args[field].(string),
		}
	}
	return pgtype.Text{
		Valid:  false,
		String: "",
	}
}

func UuidToPgTypeUuid(ID uuid.UUID) pgtype.UUID {
	return pgtype.UUID{
		Bytes: ID,
		Valid: ID != uuid.Nil,
	}
}
