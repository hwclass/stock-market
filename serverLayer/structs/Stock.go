package Stock

type Stock struct {
  Id int64 `db:"id" json:"id"`
  Firm string `db:"firm" json:"firm"`
  Data string `db:"data" json:"data"`
}