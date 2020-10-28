// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package hotelmodel

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

func easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels(in *jlexer.Lexer, out *SearchString) {
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
		case "pattern":
			out.Pattern = string(in.String())
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
func easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels(out *jwriter.Writer, in SearchString) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"pattern\":"
		out.RawString(prefix[1:])
		out.String(string(in.Pattern))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchString) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchString) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchString) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchString) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels(l, v)
}
func easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels1(in *jlexer.Lexer, out *SearchData) {
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
		case "hotels":
			if in.IsNull() {
				in.Skip()
				out.Hotels = nil
			} else {
				in.Delim('[')
				if out.Hotels == nil {
					if !in.IsDelim(']') {
						out.Hotels = make([]Hotel, 0, 0)
					} else {
						out.Hotels = []Hotel{}
					}
				} else {
					out.Hotels = (out.Hotels)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Hotel
					(v1).UnmarshalEasyJSON(in)
					out.Hotels = append(out.Hotels, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "pag_info":
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
func easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels1(out *jwriter.Writer, in SearchData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"hotels\":"
		out.RawString(prefix[1:])
		if in.Hotels == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Hotels {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"pag_info\":"
		out.RawString(prefix)
		(in.PagInfo).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels1(l, v)
}
func easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels2(in *jlexer.Lexer, out *HotelPreview) {
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
		case "hotel_id":
			out.HotelID = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "image":
			out.Image = string(in.String())
		case "location":
			out.Location = string(in.String())
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
func easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels2(out *jwriter.Writer, in HotelPreview) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"hotel_id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.HotelID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"image\":"
		out.RawString(prefix)
		out.String(string(in.Image))
	}
	{
		const prefix string = ",\"location\":"
		out.RawString(prefix)
		out.String(string(in.Location))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HotelPreview) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HotelPreview) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HotelPreview) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HotelPreview) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels2(l, v)
}
func easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels3(in *jlexer.Lexer, out *HotelData) {
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
		case "hotel":
			(out.Hotel).UnmarshalEasyJSON(in)
		case "rate":
			out.CurrRate = int(in.Int())
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
func easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels3(out *jwriter.Writer, in HotelData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"hotel\":"
		out.RawString(prefix[1:])
		(in.Hotel).MarshalEasyJSON(out)
	}
	if in.CurrRate != 0 {
		const prefix string = ",\"rate\":"
		out.RawString(prefix)
		out.Int(int(in.CurrRate))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HotelData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HotelData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HotelData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HotelData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels3(l, v)
}
func easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels4(in *jlexer.Lexer, out *Hotel) {
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
		case "hotel_id":
			out.HotelID = int(in.Int())
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "image":
			out.Image = string(in.String())
		case "location":
			out.Location = string(in.String())
		case "rating":
			out.Rating = float64(in.Float64())
		case "photos":
			if in.IsNull() {
				in.Skip()
				out.Photos = nil
			} else {
				in.Delim('[')
				if out.Photos == nil {
					if !in.IsDelim(']') {
						out.Photos = make([]string, 0, 4)
					} else {
						out.Photos = []string{}
					}
				} else {
					out.Photos = (out.Photos)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Photos = append(out.Photos, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "comm_count":
			out.CommCount = int(in.Int())
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
func easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels4(out *jwriter.Writer, in Hotel) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"hotel_id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.HotelID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"image\":"
		out.RawString(prefix)
		out.String(string(in.Image))
	}
	{
		const prefix string = ",\"location\":"
		out.RawString(prefix)
		out.String(string(in.Location))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.Float64(float64(in.Rating))
	}
	if len(in.Photos) != 0 {
		const prefix string = ",\"photos\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v5, v6 := range in.Photos {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"comm_count\":"
		out.RawString(prefix)
		out.Int(int(in.CommCount))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Hotel) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Hotel) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Hotel) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Hotel) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels4(l, v)
}
func easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels5(in *jlexer.Lexer, out *FilterData) {
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
		case "Rating":
			out.Rating = float64(in.Float64())
		case "ID":
			out.ID = string(in.String())
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
func easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels5(out *jwriter.Writer, in FilterData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Rating\":"
		out.RawString(prefix[1:])
		out.Float64(float64(in.Rating))
	}
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FilterData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FilterData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FilterData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FilterData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels5(l, v)
}
func easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels6(in *jlexer.Lexer, out *Cursor) {
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
		case "nextcursor":
			out.NextCursor = string(in.String())
		case "prevcursor":
			out.PrevCursor = string(in.String())
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
func easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels6(out *jwriter.Writer, in Cursor) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nextcursor\":"
		out.RawString(prefix[1:])
		out.String(string(in.NextCursor))
	}
	{
		const prefix string = ",\"prevcursor\":"
		out.RawString(prefix)
		out.String(string(in.PrevCursor))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Cursor) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Cursor) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Cursor) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Cursor) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsInternalAppHotelsModels6(l, v)
}
