// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetDrawing

import (
	"encoding/xml"

	"baliance.com/gooxml"
)

type EG_ObjectChoicesChoice struct {
	Sp           *CT_Shape
	GrpSp        *CT_GroupShape
	GraphicFrame *CT_GraphicalObjectFrame
	CxnSp        *CT_Connector
	Pic          *CT_Picture
	ContentPart  *CT_Rel
}

func NewEG_ObjectChoicesChoice() *EG_ObjectChoicesChoice {
	ret := &EG_ObjectChoicesChoice{}
	return ret
}

func (m *EG_ObjectChoicesChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Sp != nil {
		sesp := xml.StartElement{Name: xml.Name{Local: "xdr:sp"}}
		e.EncodeElement(m.Sp, sesp)
	}
	if m.GrpSp != nil {
		segrpSp := xml.StartElement{Name: xml.Name{Local: "xdr:grpSp"}}
		e.EncodeElement(m.GrpSp, segrpSp)
	}
	if m.GraphicFrame != nil {
		segraphicFrame := xml.StartElement{Name: xml.Name{Local: "xdr:graphicFrame"}}
		e.EncodeElement(m.GraphicFrame, segraphicFrame)
	}
	if m.CxnSp != nil {
		secxnSp := xml.StartElement{Name: xml.Name{Local: "xdr:cxnSp"}}
		e.EncodeElement(m.CxnSp, secxnSp)
	}
	if m.Pic != nil {
		sepic := xml.StartElement{Name: xml.Name{Local: "xdr:pic"}}
		e.EncodeElement(m.Pic, sepic)
	}
	if m.ContentPart != nil {
		secontentPart := xml.StartElement{Name: xml.Name{Local: "xdr:contentPart"}}
		e.EncodeElement(m.ContentPart, secontentPart)
	}
	return nil
}

func (m *EG_ObjectChoicesChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_ObjectChoicesChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "sp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "sp"}:
				m.Sp = NewCT_Shape()
				if err := d.DecodeElement(m.Sp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "grpSp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "grpSp"}:
				m.GrpSp = NewCT_GroupShape()
				if err := d.DecodeElement(m.GrpSp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "graphicFrame"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "graphicFrame"}:
				m.GraphicFrame = NewCT_GraphicalObjectFrame()
				if err := d.DecodeElement(m.GraphicFrame, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "cxnSp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "cxnSp"}:
				m.CxnSp = NewCT_Connector()
				if err := d.DecodeElement(m.CxnSp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "pic"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "pic"}:
				m.Pic = NewCT_Picture()
				if err := d.DecodeElement(m.Pic, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "contentPart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "contentPart"}:
				m.ContentPart = NewCT_Rel()
				if err := d.DecodeElement(m.ContentPart, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on EG_ObjectChoicesChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_ObjectChoicesChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ObjectChoicesChoice and its children
func (m *EG_ObjectChoicesChoice) Validate() error {
	return m.ValidateWithPath("EG_ObjectChoicesChoice")
}

// ValidateWithPath validates the EG_ObjectChoicesChoice and its children, prefixing error messages with path
func (m *EG_ObjectChoicesChoice) ValidateWithPath(path string) error {
	if m.Sp != nil {
		if err := m.Sp.ValidateWithPath(path + "/Sp"); err != nil {
			return err
		}
	}
	if m.GrpSp != nil {
		if err := m.GrpSp.ValidateWithPath(path + "/GrpSp"); err != nil {
			return err
		}
	}
	if m.GraphicFrame != nil {
		if err := m.GraphicFrame.ValidateWithPath(path + "/GraphicFrame"); err != nil {
			return err
		}
	}
	if m.CxnSp != nil {
		if err := m.CxnSp.ValidateWithPath(path + "/CxnSp"); err != nil {
			return err
		}
	}
	if m.Pic != nil {
		if err := m.Pic.ValidateWithPath(path + "/Pic"); err != nil {
			return err
		}
	}
	if m.ContentPart != nil {
		if err := m.ContentPart.ValidateWithPath(path + "/ContentPart"); err != nil {
			return err
		}
	}
	return nil
}
