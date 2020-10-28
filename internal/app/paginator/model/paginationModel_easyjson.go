// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package paginationModel

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson47583779DecodeGithubComGoParkMailRu20202JMickhsInternalAppPaginatorModel(in *jlexer.Lexer, out *PaginationModel) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "list":
			if m, ok := out.List.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.List.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.List = in.Interface()
			}
		case "Pag_info":
			(out.PagInfo).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson47583779EncodeGithubComGoParkMailRu20202JMickhsInternalAppPaginatorModel(out *jwriter.Writer, in PaginationModel) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"list\":"
		out.RawString(prefix[1:])
		if m, ok := in.List.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.List.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.List))
		}
	}
	{
		const prefix string = ",\"Pag_info\":"
		out.RawString(prefix)
		(in.PagInfo).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PaginationModel) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson47583779EncodeGithubComGoParkMailRu20202JMickhsInternalAppPaginatorModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PaginationModel) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson47583779EncodeGithubComGoParkMailRu20202JMickhsInternalAppPaginatorModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PaginationModel) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson47583779DecodeGithubComGoParkMailRu20202JMickhsInternalAppPaginatorModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PaginationModel) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson47583779DecodeGithubComGoParkMailRu20202JMickhsInternalAppPaginatorModel(l, v)
}
func easyjson47583779DecodeGithubComGoParkMailRu20202JMickhsInternalAppPaginatorModel1(in *jlexer.Lexer, out *PaginationInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "page_num":
			out.PageNum = int(in.Int())
		case "has_next":
			out.HasNext = bool(in.Bool())
		case "has_prev":
			out.HasPrev = bool(in.Bool())
		case "num_pages":
			out.NumPages = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson47583779EncodeGithubComGoParkMailRu20202JMickhsInternalAppPaginatorModel1(out *jwriter.Writer, in PaginationInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"page_num\":"
		out.RawString(prefix[1:])
		out.Int(int(in.PageNum))
	}
	{
		const prefix string = ",\"has_next\":"
		out.RawString(prefix)
		out.Bool(bool(in.HasNext))
	}
	{
		const prefix string = ",\"has_prev\":"
		out.RawString(prefix)
		out.Bool(bool(in.HasPrev))
	}
	{
		const prefix string = ",\"num_pages\":"
		out.RawString(prefix)
		out.Int(int(in.NumPages))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PaginationInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson47583779EncodeGithubComGoParkMailRu20202JMickhsInternalAppPaginatorModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PaginationInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson47583779EncodeGithubComGoParkMailRu20202JMickhsInternalAppPaginatorModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PaginationInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson47583779DecodeGithubComGoParkMailRu20202JMickhsInternalAppPaginatorModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PaginationInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson47583779DecodeGithubComGoParkMailRu20202JMickhsInternalAppPaginatorModel1(l, v)
}