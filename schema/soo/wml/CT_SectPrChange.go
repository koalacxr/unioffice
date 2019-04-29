// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	"baliance.com/gooxml"
)

type CT_SectPrChange struct {
	AuthorAttr string
	DateAttr   *time.Time
	// Annotation Identifier
	IdAttr int64
	SectPr *CT_SectPrBase
}

func NewCT_SectPrChange() *CT_SectPrChange {
	ret := &CT_SectPrChange{}
	return ret
}

func (m *CT_SectPrChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:author"},
		Value: fmt.Sprintf("%v", m.AuthorAttr)})
	if m.DateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:date"},
			Value: fmt.Sprintf("%v", *m.DateAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	if m.SectPr != nil {
		sesectPr := xml.StartElement{Name: xml.Name{Local: "w:sectPr"}}
		e.EncodeElement(m.SectPr, sesectPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SectPrChange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "author" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AuthorAttr = parsed
			continue
		}
		if attr.Name.Local == "date" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.DateAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
			continue
		}
	}
lCT_SectPrChange:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "sectPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "sectPr"}:
				m.SectPr = NewCT_SectPrBase()
				if err := d.DecodeElement(m.SectPr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_SectPrChange %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SectPrChange
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SectPrChange and its children
func (m *CT_SectPrChange) Validate() error {
	return m.ValidateWithPath("CT_SectPrChange")
}

// ValidateWithPath validates the CT_SectPrChange and its children, prefixing error messages with path
func (m *CT_SectPrChange) ValidateWithPath(path string) error {
	if m.SectPr != nil {
		if err := m.SectPr.ValidateWithPath(path + "/SectPr"); err != nil {
			return err
		}
	}
	return nil
}
