// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package http_models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	models "patreon/internal/app/models"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels(in *jlexer.Lexer, out *SubscribeRequest) {
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
		case "pay_token":
			out.Token = string(in.String())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels(out *jwriter.Writer, in SubscribeRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"pay_token\":"
		out.RawString(prefix[1:])
		out.String(string(in.Token))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SubscribeRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SubscribeRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SubscribeRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SubscribeRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels(l, v)
}
func easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels1(in *jlexer.Lexer, out *RequestText) {
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
		case "text":
			out.Text = string(in.String())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels1(out *jwriter.Writer, in RequestText) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix[1:])
		out.String(string(in.Text))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestText) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestText) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestText) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestText) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels1(l, v)
}
func easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels2(in *jlexer.Lexer, out *RequestRegistration) {
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
		case "login":
			out.Login = string(in.String())
		case "nickname":
			out.Nickname = string(in.String())
		case "password":
			out.Password = string(in.String())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels2(out *jwriter.Writer, in RequestRegistration) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"login\":"
		out.RawString(prefix[1:])
		out.String(string(in.Login))
	}
	{
		const prefix string = ",\"nickname\":"
		out.RawString(prefix)
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestRegistration) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestRegistration) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestRegistration) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestRegistration) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels2(l, v)
}
func easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels3(in *jlexer.Lexer, out *RequestPosts) {
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
		case "title":
			out.Title = string(in.String())
		case "awards_id":
			out.AwardsId = int64(in.Int64())
		case "description":
			out.Description = string(in.String())
		case "is_draft":
			out.IsDraft = bool(in.Bool())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels3(out *jwriter.Writer, in RequestPosts) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Title != "" {
		const prefix string = ",\"title\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	if in.AwardsId != 0 {
		const prefix string = ",\"awards_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.AwardsId))
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if in.IsDraft {
		const prefix string = ",\"is_draft\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsDraft))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestPosts) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestPosts) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestPosts) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestPosts) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels3(l, v)
}
func easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels4(in *jlexer.Lexer, out *RequestLogin) {
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
		case "login":
			out.Login = string(in.String())
		case "password":
			out.Password = string(in.String())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels4(out *jwriter.Writer, in RequestLogin) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"login\":"
		out.RawString(prefix[1:])
		out.String(string(in.Login))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestLogin) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestLogin) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestLogin) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestLogin) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels4(l, v)
}
func easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels5(in *jlexer.Lexer, out *RequestCreator) {
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
		case "category":
			out.Category = string(in.String())
		case "description":
			out.Description = string(in.String())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels5(out *jwriter.Writer, in RequestCreator) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"category\":"
		out.RawString(prefix[1:])
		out.String(string(in.Category))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestCreator) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestCreator) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestCreator) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestCreator) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels5(l, v)
}
func easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels6(in *jlexer.Lexer, out *RequestComment) {
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
		case "body":
			out.Body = string(in.String())
		case "as_creator":
			out.AsCreator = bool(in.Bool())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels6(out *jwriter.Writer, in RequestComment) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"body\":"
		out.RawString(prefix[1:])
		out.String(string(in.Body))
	}
	if in.AsCreator {
		const prefix string = ",\"as_creator\":"
		out.RawString(prefix)
		out.Bool(bool(in.AsCreator))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestComment) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestComment) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestComment) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestComment) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels6(l, v)
}
func easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels7(in *jlexer.Lexer, out *RequestChangePassword) {
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
		case "old":
			out.OldPassword = string(in.String())
		case "new":
			out.NewPassword = string(in.String())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels7(out *jwriter.Writer, in RequestChangePassword) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"old\":"
		out.RawString(prefix[1:])
		out.String(string(in.OldPassword))
	}
	{
		const prefix string = ",\"new\":"
		out.RawString(prefix)
		out.String(string(in.NewPassword))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestChangePassword) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestChangePassword) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestChangePassword) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestChangePassword) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels7(l, v)
}
func easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels8(in *jlexer.Lexer, out *RequestChangeNickname) {
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
		case "old":
			out.OldNickname = string(in.String())
		case "new":
			out.NewNickname = string(in.String())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels8(out *jwriter.Writer, in RequestChangeNickname) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"old\":"
		out.RawString(prefix[1:])
		out.String(string(in.OldNickname))
	}
	{
		const prefix string = ",\"new\":"
		out.RawString(prefix)
		out.String(string(in.NewNickname))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestChangeNickname) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestChangeNickname) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestChangeNickname) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestChangeNickname) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels8(l, v)
}
func easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels9(in *jlexer.Lexer, out *RequestAwards) {
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
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "price":
			out.Price = int64(in.Int64())
		case "color":
			(out.Color).UnmarshalEasyJSON(in)
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels9(out *jwriter.Writer, in RequestAwards) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix)
		out.Int64(int64(in.Price))
	}
	if true {
		const prefix string = ",\"color\":"
		out.RawString(prefix)
		(in.Color).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestAwards) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestAwards) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestAwards) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestAwards) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels9(l, v)
}
func easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels10(in *jlexer.Lexer, out *RequestAttaches) {
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
		case "attaches":
			if in.IsNull() {
				in.Skip()
				out.Attaches = nil
			} else {
				in.Delim('[')
				if out.Attaches == nil {
					if !in.IsDelim(']') {
						out.Attaches = make([]RequestAttach, 0, 1)
					} else {
						out.Attaches = []RequestAttach{}
					}
				} else {
					out.Attaches = (out.Attaches)[:0]
				}
				for !in.IsDelim(']') {
					var v1 RequestAttach
					(v1).UnmarshalEasyJSON(in)
					out.Attaches = append(out.Attaches, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels10(out *jwriter.Writer, in RequestAttaches) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"attaches\":"
		out.RawString(prefix[1:])
		if in.Attaches == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Attaches {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestAttaches) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestAttaches) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestAttaches) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestAttaches) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels10(l, v)
}
func easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels11(in *jlexer.Lexer, out *RequestAttach) {
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
		case "type":
			out.Type = models.DataType(in.String())
		case "value":
			out.Value = string(in.String())
		case "id":
			out.Id = int64(in.Int64())
		case "status":
			out.Status = string(in.String())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels11(out *jwriter.Writer, in RequestAttach) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix[1:])
		out.String(string(in.Type))
	}
	if in.Value != "" {
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		out.String(string(in.Value))
	}
	if in.Id != 0 {
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Int64(int64(in.Id))
	}
	if in.Status != "" {
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestAttach) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestAttach) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestAttach) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestAttach) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels11(l, v)
}
func easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels12(in *jlexer.Lexer, out *Color) {
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
		case "red":
			out.R = uint8(in.Uint8())
		case "green":
			out.G = uint8(in.Uint8())
		case "blue":
			out.B = uint8(in.Uint8())
		case "alpha":
			out.A = uint8(in.Uint8())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels12(out *jwriter.Writer, in Color) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"red\":"
		out.RawString(prefix[1:])
		out.Uint8(uint8(in.R))
	}
	{
		const prefix string = ",\"green\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.G))
	}
	{
		const prefix string = ",\"blue\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.B))
	}
	{
		const prefix string = ",\"alpha\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.A))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Color) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Color) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7df0efccEncodePatreonInternalAppDeliveryHttpModels12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Color) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Color) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7df0efccDecodePatreonInternalAppDeliveryHttpModels12(l, v)
}
