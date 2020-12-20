// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package chat_model

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

func easyjson5b4733c0DecodeGithubComGoParkMailRu20202JMickhsMainInternalAppChatModels(in *jlexer.Lexer, out *Message) {
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
		case "OwnerID":
			out.OwnerID = string(in.String())
		case "Room":
			out.Room = string(in.String())
		case "Message":
			out.Message = string(in.String())
		case "Moderator":
			out.Moderator = bool(in.Bool())
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
func easyjson5b4733c0EncodeGithubComGoParkMailRu20202JMickhsMainInternalAppChatModels(out *jwriter.Writer, in Message) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"OwnerID\":"
		out.RawString(prefix[1:])
		out.String(string(in.OwnerID))
	}
	{
		const prefix string = ",\"Room\":"
		out.RawString(prefix)
		out.String(string(in.Room))
	}
	{
		const prefix string = ",\"Message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	{
		const prefix string = ",\"Moderator\":"
		out.RawString(prefix)
		out.Bool(bool(in.Moderator))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Message) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5b4733c0EncodeGithubComGoParkMailRu20202JMickhsMainInternalAppChatModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Message) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5b4733c0EncodeGithubComGoParkMailRu20202JMickhsMainInternalAppChatModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Message) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5b4733c0DecodeGithubComGoParkMailRu20202JMickhsMainInternalAppChatModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Message) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5b4733c0DecodeGithubComGoParkMailRu20202JMickhsMainInternalAppChatModels(l, v)
}
