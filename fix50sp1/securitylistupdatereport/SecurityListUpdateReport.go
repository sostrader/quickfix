//Package securitylistupdatereport msg type = BK.
package securitylistupdatereport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix/enum"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a SecurityListUpdateReport wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//SecurityReportID is a non-required field for SecurityListUpdateReport.
func (m Message) SecurityReportID() (*field.SecurityReportIDField, quickfix.MessageRejectError) {
	f := &field.SecurityReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReportID reads a SecurityReportID from SecurityListUpdateReport.
func (m Message) GetSecurityReportID(f *field.SecurityReportIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityReqID is a non-required field for SecurityListUpdateReport.
func (m Message) SecurityReqID() (*field.SecurityReqIDField, quickfix.MessageRejectError) {
	f := &field.SecurityReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from SecurityListUpdateReport.
func (m Message) GetSecurityReqID(f *field.SecurityReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseID is a non-required field for SecurityListUpdateReport.
func (m Message) SecurityResponseID() (*field.SecurityResponseIDField, quickfix.MessageRejectError) {
	f := &field.SecurityResponseIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseID reads a SecurityResponseID from SecurityListUpdateReport.
func (m Message) GetSecurityResponseID(f *field.SecurityResponseIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityRequestResult is a non-required field for SecurityListUpdateReport.
func (m Message) SecurityRequestResult() (*field.SecurityRequestResultField, quickfix.MessageRejectError) {
	f := &field.SecurityRequestResultField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityRequestResult reads a SecurityRequestResult from SecurityListUpdateReport.
func (m Message) GetSecurityRequestResult(f *field.SecurityRequestResultField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoRelatedSym is a non-required field for SecurityListUpdateReport.
func (m Message) TotNoRelatedSym() (*field.TotNoRelatedSymField, quickfix.MessageRejectError) {
	f := &field.TotNoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoRelatedSym reads a TotNoRelatedSym from SecurityListUpdateReport.
func (m Message) GetTotNoRelatedSym(f *field.TotNoRelatedSymField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for SecurityListUpdateReport.
func (m Message) ClearingBusinessDate() (*field.ClearingBusinessDateField, quickfix.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from SecurityListUpdateReport.
func (m Message) GetClearingBusinessDate(f *field.ClearingBusinessDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityUpdateAction is a non-required field for SecurityListUpdateReport.
func (m Message) SecurityUpdateAction() (*field.SecurityUpdateActionField, quickfix.MessageRejectError) {
	f := &field.SecurityUpdateActionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityUpdateAction reads a SecurityUpdateAction from SecurityListUpdateReport.
func (m Message) GetSecurityUpdateAction(f *field.SecurityUpdateActionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CorporateAction is a non-required field for SecurityListUpdateReport.
func (m Message) CorporateAction() (*field.CorporateActionField, quickfix.MessageRejectError) {
	f := &field.CorporateActionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCorporateAction reads a CorporateAction from SecurityListUpdateReport.
func (m Message) GetCorporateAction(f *field.CorporateActionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for SecurityListUpdateReport.
func (m Message) LastFragment() (*field.LastFragmentField, quickfix.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from SecurityListUpdateReport.
func (m Message) GetLastFragment(f *field.LastFragmentField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for SecurityListUpdateReport.
func (m Message) NoRelatedSym() (*field.NoRelatedSymField, quickfix.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from SecurityListUpdateReport.
func (m Message) GetNoRelatedSym(f *field.NoRelatedSymField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for SecurityListUpdateReport.
func (m Message) MarketID() (*field.MarketIDField, quickfix.MessageRejectError) {
	f := &field.MarketIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from SecurityListUpdateReport.
func (m Message) GetMarketID(f *field.MarketIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for SecurityListUpdateReport.
func (m Message) MarketSegmentID() (*field.MarketSegmentIDField, quickfix.MessageRejectError) {
	f := &field.MarketSegmentIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from SecurityListUpdateReport.
func (m Message) GetMarketSegmentID(f *field.MarketSegmentIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for SecurityListUpdateReport.
func (m Message) ApplID() (*field.ApplIDField, quickfix.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from SecurityListUpdateReport.
func (m Message) GetApplID(f *field.ApplIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for SecurityListUpdateReport.
func (m Message) ApplSeqNum() (*field.ApplSeqNumField, quickfix.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from SecurityListUpdateReport.
func (m Message) GetApplSeqNum(f *field.ApplSeqNumField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for SecurityListUpdateReport.
func (m Message) ApplLastSeqNum() (*field.ApplLastSeqNumField, quickfix.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from SecurityListUpdateReport.
func (m Message) GetApplLastSeqNum(f *field.ApplLastSeqNumField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for SecurityListUpdateReport.
func (m Message) ApplResendFlag() (*field.ApplResendFlagField, quickfix.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from SecurityListUpdateReport.
func (m Message) GetApplResendFlag(f *field.ApplResendFlagField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for SecurityListUpdateReport.
func New() Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.NewMsgType("BK"))
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "BK", r
}
