// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vml

import (
	"encoding/xml"

	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/soo/wml"
)

type Textbox struct {
	CT_Textbox
}

func NewTextbox() *Textbox {
	ret := &Textbox{}
	ret.CT_Textbox = *NewCT_Textbox()
	return ret
}

func (m *Textbox) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_Textbox.MarshalXML(e, start)
}

func (m *Textbox) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Textbox = *NewCT_Textbox()
	for _, attr := range start.Attr {
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "singleclick" {
			m.SingleclickAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "insetmode" {
			m.InsetmodeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "inset" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.InsetAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "style" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StyleAttr = &parsed
			continue
		}
	}
lTextbox:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "txbxContent"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "txbxContent"}:
				m.TxbxContent = wml.NewTxbxContent()
				if err := d.DecodeElement(m.TxbxContent, &el); err != nil {
					return err
				}
			default:
				if anyEl, err := gooxml.CreateElement(el); err != nil {
					return err
				} else {
					if err := d.DecodeElement(anyEl, &el); err != nil {
						return err
					}
					m.Any = anyEl
				}
			}
		case xml.EndElement:
			break lTextbox
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Textbox and its children
func (m *Textbox) Validate() error {
	return m.ValidateWithPath("Textbox")
}

// ValidateWithPath validates the Textbox and its children, prefixing error messages with path
func (m *Textbox) ValidateWithPath(path string) error {
	if err := m.CT_Textbox.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
