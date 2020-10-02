// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package benchmarks

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

func easyjson890029d8DecodeGithubComYohanson555GojayBenchmarks(in *jlexer.Lexer, out *SmallPayload) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "St":
			out.St = int(in.Int())
		case "Sid":
			out.Sid = int(in.Int())
		case "Tt":
			out.Tt = string(in.String())
		case "Gr":
			out.Gr = int(in.Int())
		case "Uuid":
			out.Uuid = string(in.String())
		case "Ip":
			out.Ip = string(in.String())
		case "Ua":
			out.Ua = string(in.String())
		case "Tz":
			out.Tz = int(in.Int())
		case "V":
			out.V = int(in.Int())
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
func easyjson890029d8EncodeGithubComYohanson555GojayBenchmarks(out *jwriter.Writer, in SmallPayload) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"St\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.St))
	}
	{
		const prefix string = ",\"Sid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Sid))
	}
	{
		const prefix string = ",\"Tt\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Tt))
	}
	{
		const prefix string = ",\"Gr\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Gr))
	}
	{
		const prefix string = ",\"Uuid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Uuid))
	}
	{
		const prefix string = ",\"Ip\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Ip))
	}
	{
		const prefix string = ",\"Ua\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Ua))
	}
	{
		const prefix string = ",\"Tz\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Tz))
	}
	{
		const prefix string = ",\"V\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.V))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SmallPayload) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson890029d8EncodeGithubComYohanson555GojayBenchmarks(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SmallPayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson890029d8EncodeGithubComYohanson555GojayBenchmarks(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SmallPayload) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson890029d8DecodeGithubComYohanson555GojayBenchmarks(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SmallPayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson890029d8DecodeGithubComYohanson555GojayBenchmarks(l, v)
}
