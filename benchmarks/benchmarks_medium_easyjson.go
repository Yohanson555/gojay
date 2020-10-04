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

func easyjsonB0500db0DecodeGithubComfrancoispqtGojayBenchmarks(in *jlexer.Lexer, out *MediumPayload) {
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
		case "person":
			if in.IsNull() {
				in.Skip()
				out.Person = nil
			} else {
				if out.Person == nil {
					out.Person = new(CBPerson)
				}
				(*out.Person).UnmarshalEasyJSON(in)
			}
		case "company":
			out.Company = string(in.String())
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
func easyjsonB0500db0EncodeGithubComfrancoispqtGojayBenchmarks(out *jwriter.Writer, in MediumPayload) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"person\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Person == nil {
			out.RawString("null")
		} else {
			(*in.Person).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"company\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Company))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MediumPayload) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB0500db0EncodeGithubComfrancoispqtGojayBenchmarks(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MediumPayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB0500db0EncodeGithubComfrancoispqtGojayBenchmarks(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MediumPayload) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB0500db0DecodeGithubComfrancoispqtGojayBenchmarks(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MediumPayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB0500db0DecodeGithubComfrancoispqtGojayBenchmarks(l, v)
}
func easyjsonB0500db0DecodeGithubComfrancoispqtGojayBenchmarks1(in *jlexer.Lexer, out *CBPerson) {
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
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(CBName)
				}
				(*out.Name).UnmarshalEasyJSON(in)
			}
		case "github":
			if in.IsNull() {
				in.Skip()
				out.Github = nil
			} else {
				if out.Github == nil {
					out.Github = new(CBGithub)
				}
				(*out.Github).UnmarshalEasyJSON(in)
			}
		case "Gravatar":
			if in.IsNull() {
				in.Skip()
				out.Gravatar = nil
			} else {
				if out.Gravatar == nil {
					out.Gravatar = new(CBGravatar)
				}
				(*out.Gravatar).UnmarshalEasyJSON(in)
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
func easyjsonB0500db0EncodeGithubComfrancoispqtGojayBenchmarks1(out *jwriter.Writer, in CBPerson) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Name == nil {
			out.RawString("null")
		} else {
			(*in.Name).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"github\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Github == nil {
			out.RawString("null")
		} else {
			(*in.Github).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"Gravatar\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Gravatar == nil {
			out.RawString("null")
		} else {
			(*in.Gravatar).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CBPerson) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB0500db0EncodeGithubComfrancoispqtGojayBenchmarks1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CBPerson) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB0500db0EncodeGithubComfrancoispqtGojayBenchmarks1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CBPerson) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB0500db0DecodeGithubComfrancoispqtGojayBenchmarks1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CBPerson) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB0500db0DecodeGithubComfrancoispqtGojayBenchmarks1(l, v)
}
func easyjsonB0500db0DecodeGithubComfrancoispqtGojayBenchmarks2(in *jlexer.Lexer, out *CBName) {
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
		case "fullName":
			out.FullName = string(in.String())
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
func easyjsonB0500db0EncodeGithubComfrancoispqtGojayBenchmarks2(out *jwriter.Writer, in CBName) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"fullName\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FullName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CBName) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB0500db0EncodeGithubComfrancoispqtGojayBenchmarks2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CBName) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB0500db0EncodeGithubComfrancoispqtGojayBenchmarks2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CBName) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB0500db0DecodeGithubComfrancoispqtGojayBenchmarks2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CBName) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB0500db0DecodeGithubComfrancoispqtGojayBenchmarks2(l, v)
}
func easyjsonB0500db0DecodeGithubComfrancoispqtGojayBenchmarks3(in *jlexer.Lexer, out *CBGravatar) {
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
		case "Avatars":
			if in.IsNull() {
				in.Skip()
				out.Avatars = nil
			} else {
				in.Delim('[')
				if out.Avatars == nil {
					if !in.IsDelim(']') {
						out.Avatars = make(Avatars, 0, 8)
					} else {
						out.Avatars = Avatars{}
					}
				} else {
					out.Avatars = (out.Avatars)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *CBAvatar
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(CBAvatar)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Avatars = append(out.Avatars, v1)
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
func easyjsonB0500db0EncodeGithubComfrancoispqtGojayBenchmarks3(out *jwriter.Writer, in CBGravatar) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Avatars\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Avatars == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Avatars {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CBGravatar) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB0500db0EncodeGithubComfrancoispqtGojayBenchmarks3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CBGravatar) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB0500db0EncodeGithubComfrancoispqtGojayBenchmarks3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CBGravatar) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB0500db0DecodeGithubComfrancoispqtGojayBenchmarks3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CBGravatar) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB0500db0DecodeGithubComfrancoispqtGojayBenchmarks3(l, v)
}
func easyjsonB0500db0DecodeGithubComfrancoispqtGojayBenchmarks4(in *jlexer.Lexer, out *CBGithub) {
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
		case "Followers":
			out.Followers = int(in.Int())
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
func easyjsonB0500db0EncodeGithubComfrancoispqtGojayBenchmarks4(out *jwriter.Writer, in CBGithub) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Followers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Followers))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CBGithub) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB0500db0EncodeGithubComfrancoispqtGojayBenchmarks4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CBGithub) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB0500db0EncodeGithubComfrancoispqtGojayBenchmarks4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CBGithub) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB0500db0DecodeGithubComfrancoispqtGojayBenchmarks4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CBGithub) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB0500db0DecodeGithubComfrancoispqtGojayBenchmarks4(l, v)
}
func easyjsonB0500db0DecodeGithubComfrancoispqtGojayBenchmarks5(in *jlexer.Lexer, out *CBAvatar) {
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
		case "Url":
			out.Url = string(in.String())
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
func easyjsonB0500db0EncodeGithubComfrancoispqtGojayBenchmarks5(out *jwriter.Writer, in CBAvatar) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Url))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CBAvatar) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB0500db0EncodeGithubComfrancoispqtGojayBenchmarks5(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CBAvatar) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB0500db0DecodeGithubComfrancoispqtGojayBenchmarks5(l, v)
}
