package paginationModel

type PaginationModel struct {
	List    interface{}    `json:"list"`
	PagInfo PaginationInfo `json:"Pag_info"`
}
type PaginationInfo struct {
	PageNum  int  `json:"page_num"`
	HasNext  bool `json:"has_next"`
	HasPrev  bool `json:"has_prev"`
	NumPages int  `json:"num_pages"`
}
