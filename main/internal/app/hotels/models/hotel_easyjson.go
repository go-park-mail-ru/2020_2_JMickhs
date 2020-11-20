// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package hotelmodel

import (
	json "encoding/json"
	models "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/comment/models"
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

func easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels(in *jlexer.Lexer, out *SearchString) {
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
func easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels(out *jwriter.Writer, in SearchString) {
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
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchString) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchString) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchString) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels(l, v)
}
func easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels1(in *jlexer.Lexer, out *SearchData) {
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
func easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels1(out *jwriter.Writer, in SearchData) {
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
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels1(l, v)
}
func easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels2(in *jlexer.Lexer, out *HotelsPreview) {
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
		case "hotels_preview":
			if in.IsNull() {
				in.Skip()
				out.Hotels = nil
			} else {
				in.Delim('[')
				if out.Hotels == nil {
					if !in.IsDelim(']') {
						out.Hotels = make([]HotelPreview, 0, 1)
					} else {
						out.Hotels = []HotelPreview{}
					}
				} else {
					out.Hotels = (out.Hotels)[:0]
				}
				for !in.IsDelim(']') {
					var v4 HotelPreview
					(v4).UnmarshalEasyJSON(in)
					out.Hotels = append(out.Hotels, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels2(out *jwriter.Writer, in HotelsPreview) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"hotels_preview\":"
		out.RawString(prefix[1:])
		if in.Hotels == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Hotels {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HotelsPreview) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HotelsPreview) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HotelsPreview) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HotelsPreview) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels2(l, v)
}
func easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels3(in *jlexer.Lexer, out *Hotels) {
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
					var v7 Hotel
					(v7).UnmarshalEasyJSON(in)
					out.Hotels = append(out.Hotels, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels3(out *jwriter.Writer, in Hotels) {
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
			for v8, v9 := range in.Hotels {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Hotels) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Hotels) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Hotels) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Hotels) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels3(l, v)
}
func easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels4(in *jlexer.Lexer, out *HotelPreview) {
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
func easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels4(out *jwriter.Writer, in HotelPreview) {
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
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HotelPreview) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HotelPreview) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HotelPreview) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels4(l, v)
}
func easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels5(in *jlexer.Lexer, out *HotelFiltering) {
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
		case "RatingFilterStartNumber":
			out.RatingFilterStartNumber = string(in.String())
		case "CommentsFilterStartNumber":
			out.CommentsFilterStartNumber = string(in.String())
		case "Longitude":
			out.Longitude = string(in.String())
		case "Latitude":
			out.Latitude = string(in.String())
		case "Radius":
			out.Radius = string(in.String())
		case "CommCountConstraint":
			out.CommCountConstraint = string(in.String())
		case "CommCountPercent":
			out.CommCountPercent = string(in.String())
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
func easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels5(out *jwriter.Writer, in HotelFiltering) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"RatingFilterStartNumber\":"
		out.RawString(prefix[1:])
		out.String(string(in.RatingFilterStartNumber))
	}
	{
		const prefix string = ",\"CommentsFilterStartNumber\":"
		out.RawString(prefix)
		out.String(string(in.CommentsFilterStartNumber))
	}
	{
		const prefix string = ",\"Longitude\":"
		out.RawString(prefix)
		out.String(string(in.Longitude))
	}
	{
		const prefix string = ",\"Latitude\":"
		out.RawString(prefix)
		out.String(string(in.Latitude))
	}
	{
		const prefix string = ",\"Radius\":"
		out.RawString(prefix)
		out.String(string(in.Radius))
	}
	{
		const prefix string = ",\"CommCountConstraint\":"
		out.RawString(prefix)
		out.String(string(in.CommCountConstraint))
	}
	{
		const prefix string = ",\"CommCountPercent\":"
		out.RawString(prefix)
		out.String(string(in.CommCountPercent))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HotelFiltering) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HotelFiltering) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HotelFiltering) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HotelFiltering) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels5(l, v)
}
func easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels6(in *jlexer.Lexer, out *HotelData) {
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
		case "comment":
			if in.IsNull() {
				in.Skip()
				out.Comment = nil
			} else {
				if out.Comment == nil {
					out.Comment = new(models.FullCommentInfo)
				}
				(*out.Comment).UnmarshalEasyJSON(in)
			}
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
func easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels6(out *jwriter.Writer, in HotelData) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"hotel\":"
		out.RawString(prefix[1:])
		(in.Hotel).MarshalEasyJSON(out)
	}
	if in.Comment != nil {
		const prefix string = ",\"comment\":"
		out.RawString(prefix)
		(*in.Comment).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HotelData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HotelData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HotelData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HotelData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels6(l, v)
}
func easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels7(in *jlexer.Lexer, out *Hotel) {
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
		case "email":
			out.Email = string(in.String())
		case "country":
			out.Country = string(in.String())
		case "city":
			out.City = string(in.String())
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
					var v10 string
					v10 = string(in.String())
					out.Photos = append(out.Photos, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "comm_count":
			out.CommCount = int(in.Int())
		case "latitude":
			out.Latitude = float64(in.Float64())
		case "longitude":
			out.Longitude = float64(in.Float64())
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
func easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels7(out *jwriter.Writer, in Hotel) {
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
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"country\":"
		out.RawString(prefix)
		out.String(string(in.Country))
	}
	{
		const prefix string = ",\"city\":"
		out.RawString(prefix)
		out.String(string(in.City))
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
			for v11, v12 := range in.Photos {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"comm_count\":"
		out.RawString(prefix)
		out.Int(int(in.CommCount))
	}
	if in.Latitude != 0 {
		const prefix string = ",\"latitude\":"
		out.RawString(prefix)
		out.Float64(float64(in.Latitude))
	}
	if in.Longitude != 0 {
		const prefix string = ",\"longitude\":"
		out.RawString(prefix)
		out.Float64(float64(in.Longitude))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Hotel) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Hotel) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Hotel) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Hotel) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels7(l, v)
}
func easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels8(in *jlexer.Lexer, out *FilterData) {
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
func easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels8(out *jwriter.Writer, in FilterData) {
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
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FilterData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FilterData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FilterData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels8(l, v)
}
func easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels9(in *jlexer.Lexer, out *Cursor) {
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
func easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels9(out *jwriter.Writer, in Cursor) {
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
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Cursor) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD750f830EncodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Cursor) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Cursor) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD750f830DecodeGithubComGoParkMailRu20202JMickhsJMickhsMainInternalAppHotelsModels9(l, v)
}