// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"baliance.com/gooxml"
	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

const ST_TextScalePercentPattern = `0*(600|([0-5]?[0-9]?[0-9]))%`

var ST_TextScalePercentPatternRe = regexp.MustCompile(ST_TextScalePercentPattern)

const ST_CnfPattern = `[01]*`

var ST_CnfPatternRe = regexp.MustCompile(ST_CnfPattern)

func ParseUnionST_SignedTwipsMeasure(s string) (ST_SignedTwipsMeasure, error) {
	r := ST_SignedTwipsMeasure{}
	if sharedTypes.ST_UniversalMeasurePatternRe.MatchString(s) {
		r.ST_UniversalMeasure = &s
	} else {
		v, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return r, fmt.Errorf("parsing %s as int: %s", s, err)
		}
		r.Int64 = &v
	}
	return r, nil
}

func ParseUnionST_TwipsMeasure(s string) (sharedTypes.ST_TwipsMeasure, error) {
	ret := sharedTypes.ST_TwipsMeasure{}
	if sharedTypes.ST_PositiveUniversalMeasurePatternRe.MatchString(s) {
		ret.ST_PositiveUniversalMeasure = &s
	} else {
		v, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return ret, fmt.Errorf("parsing %s as uint: %s", s, err)
		}
		ret.ST_UnsignedDecimalNumber = &v
	}
	return ret, nil
}

func ParseUnionST_OnOff(s string) (sharedTypes.ST_OnOff, error) {
	return sharedTypes.ParseUnionST_OnOff(s)
}

func ParseUnionST_HexColor(s string) (ST_HexColor, error) {
	r := ST_HexColor{}
	if s == "auto" {
		r.ST_HexColorAuto = ST_HexColorAutoAuto
	} else {
		r.ST_HexColorRGB = &s
	}
	return r, nil
}

func ParseStdlibTime(s string) (time.Time, error) {
	return time.Time{}, nil
}

func ParseUnionST_DecimalNumberOrPercent(s string) (ST_DecimalNumberOrPercent, error) {
	ret := ST_DecimalNumberOrPercent{}
	if sharedTypes.ST_PercentagePatternRe.MatchString(s) {
		ret.ST_Percentage = &s
	} else {
		v, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return ret, fmt.Errorf("parsing %s as int: %s", s, err)
		}
		ret.ST_UnqualifiedPercentage = &v
	}
	return ret, nil
}

func ParseUnionST_MeasurementOrPercent(s string) (ST_MeasurementOrPercent, error) {
	r := ST_MeasurementOrPercent{}
	if sharedTypes.ST_UniversalMeasurePatternRe.MatchString(s) {
		r.ST_UniversalMeasure = &s
	} else {
		r.ST_DecimalNumberOrPercent = &ST_DecimalNumberOrPercent{}
		if sharedTypes.ST_PercentagePatternRe.MatchString(s) {
			r.ST_DecimalNumberOrPercent.ST_Percentage = &s
		} else {
			v, err := strconv.ParseInt(s, 10, 32)
			if err != nil {
				return r, fmt.Errorf("parsing %s as int: %s", s, err)
			}
			r.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = &v
		}
	}
	return r, nil
}

func ParseUnionST_HpsMeasure(s string) (ST_HpsMeasure, error) {
	r := ST_HpsMeasure{}

	if sharedTypes.ST_PositiveUniversalMeasurePatternRe.MatchString(s) {
		r.ST_PositiveUniversalMeasure = &s
	} else {
		v, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return r, fmt.Errorf("parsing %s as uint: %s", s, err)
		}
		r.ST_UnsignedDecimalNumber = &v
	}
	return r, nil
}

func ParseUnionST_SignedHpsMeasure(s string) (ST_SignedHpsMeasure, error) {
	r := ST_SignedHpsMeasure{}
	if sharedTypes.ST_UniversalMeasurePatternRe.MatchString(s) {
		r.ST_UniversalMeasure = &s
	} else {
		v, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return r, fmt.Errorf("parsing %s as int: %s", s, err)
		}
		r.Int64 = &v
	}
	return r, nil
}

func ParseUnionST_TextScale(s string) (ST_TextScale, error) {
	r := ST_TextScale{}
	if ST_TextScalePercentPatternRe.MatchString(s) {
		r.ST_TextScalePercent = &s
	} else {
		v, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return r, fmt.Errorf("parsing %s as int: %s", s, err)
		}
		r.ST_TextScaleDecimal = &v
	}
	return r, nil
}

type ST_HighlightColor byte

const (
	ST_HighlightColorUnset       ST_HighlightColor = 0
	ST_HighlightColorBlack       ST_HighlightColor = 1
	ST_HighlightColorBlue        ST_HighlightColor = 2
	ST_HighlightColorCyan        ST_HighlightColor = 3
	ST_HighlightColorGreen       ST_HighlightColor = 4
	ST_HighlightColorMagenta     ST_HighlightColor = 5
	ST_HighlightColorRed         ST_HighlightColor = 6
	ST_HighlightColorYellow      ST_HighlightColor = 7
	ST_HighlightColorWhite       ST_HighlightColor = 8
	ST_HighlightColorDarkBlue    ST_HighlightColor = 9
	ST_HighlightColorDarkCyan    ST_HighlightColor = 10
	ST_HighlightColorDarkGreen   ST_HighlightColor = 11
	ST_HighlightColorDarkMagenta ST_HighlightColor = 12
	ST_HighlightColorDarkRed     ST_HighlightColor = 13
	ST_HighlightColorDarkYellow  ST_HighlightColor = 14
	ST_HighlightColorDarkGray    ST_HighlightColor = 15
	ST_HighlightColorLightGray   ST_HighlightColor = 16
	ST_HighlightColorNone        ST_HighlightColor = 17
)

func (e ST_HighlightColor) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_HighlightColorUnset:
		attr.Value = ""
	case ST_HighlightColorBlack:
		attr.Value = "black"
	case ST_HighlightColorBlue:
		attr.Value = "blue"
	case ST_HighlightColorCyan:
		attr.Value = "cyan"
	case ST_HighlightColorGreen:
		attr.Value = "green"
	case ST_HighlightColorMagenta:
		attr.Value = "magenta"
	case ST_HighlightColorRed:
		attr.Value = "red"
	case ST_HighlightColorYellow:
		attr.Value = "yellow"
	case ST_HighlightColorWhite:
		attr.Value = "white"
	case ST_HighlightColorDarkBlue:
		attr.Value = "darkBlue"
	case ST_HighlightColorDarkCyan:
		attr.Value = "darkCyan"
	case ST_HighlightColorDarkGreen:
		attr.Value = "darkGreen"
	case ST_HighlightColorDarkMagenta:
		attr.Value = "darkMagenta"
	case ST_HighlightColorDarkRed:
		attr.Value = "darkRed"
	case ST_HighlightColorDarkYellow:
		attr.Value = "darkYellow"
	case ST_HighlightColorDarkGray:
		attr.Value = "darkGray"
	case ST_HighlightColorLightGray:
		attr.Value = "lightGray"
	case ST_HighlightColorNone:
		attr.Value = "none"
	}
	return attr, nil
}

func (e *ST_HighlightColor) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "black":
		*e = 1
	case "blue":
		*e = 2
	case "cyan":
		*e = 3
	case "green":
		*e = 4
	case "magenta":
		*e = 5
	case "red":
		*e = 6
	case "yellow":
		*e = 7
	case "white":
		*e = 8
	case "darkBlue":
		*e = 9
	case "darkCyan":
		*e = 10
	case "darkGreen":
		*e = 11
	case "darkMagenta":
		*e = 12
	case "darkRed":
		*e = 13
	case "darkYellow":
		*e = 14
	case "darkGray":
		*e = 15
	case "lightGray":
		*e = 16
	case "none":
		*e = 17
	}
	return nil
}

func (m ST_HighlightColor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_HighlightColor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "black":
			*m = 1
		case "blue":
			*m = 2
		case "cyan":
			*m = 3
		case "green":
			*m = 4
		case "magenta":
			*m = 5
		case "red":
			*m = 6
		case "yellow":
			*m = 7
		case "white":
			*m = 8
		case "darkBlue":
			*m = 9
		case "darkCyan":
			*m = 10
		case "darkGreen":
			*m = 11
		case "darkMagenta":
			*m = 12
		case "darkRed":
			*m = 13
		case "darkYellow":
			*m = 14
		case "darkGray":
			*m = 15
		case "lightGray":
			*m = 16
		case "none":
			*m = 17
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_HighlightColor) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "black"
	case 2:
		return "blue"
	case 3:
		return "cyan"
	case 4:
		return "green"
	case 5:
		return "magenta"
	case 6:
		return "red"
	case 7:
		return "yellow"
	case 8:
		return "white"
	case 9:
		return "darkBlue"
	case 10:
		return "darkCyan"
	case 11:
		return "darkGreen"
	case 12:
		return "darkMagenta"
	case 13:
		return "darkRed"
	case 14:
		return "darkYellow"
	case 15:
		return "darkGray"
	case 16:
		return "lightGray"
	case 17:
		return "none"
	}
	return ""
}

func (m ST_HighlightColor) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_HighlightColor) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_HexColorAuto byte

const (
	ST_HexColorAutoUnset ST_HexColorAuto = 0
	ST_HexColorAutoAuto  ST_HexColorAuto = 1
)

func (e ST_HexColorAuto) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_HexColorAutoUnset:
		attr.Value = ""
	case ST_HexColorAutoAuto:
		attr.Value = "auto"
	}
	return attr, nil
}

func (e *ST_HexColorAuto) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "auto":
		*e = 1
	}
	return nil
}

func (m ST_HexColorAuto) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_HexColorAuto) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "auto":
			*m = 1
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_HexColorAuto) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "auto"
	}
	return ""
}

func (m ST_HexColorAuto) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_HexColorAuto) ValidateWithPath(path string) error {
	switch m {
	case 0, 1:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Underline byte

const (
	ST_UnderlineUnset           ST_Underline = 0
	ST_UnderlineSingle          ST_Underline = 1
	ST_UnderlineWords           ST_Underline = 2
	ST_UnderlineDouble          ST_Underline = 3
	ST_UnderlineThick           ST_Underline = 4
	ST_UnderlineDotted          ST_Underline = 5
	ST_UnderlineDottedHeavy     ST_Underline = 6
	ST_UnderlineDash            ST_Underline = 7
	ST_UnderlineDashedHeavy     ST_Underline = 8
	ST_UnderlineDashLong        ST_Underline = 9
	ST_UnderlineDashLongHeavy   ST_Underline = 10
	ST_UnderlineDotDash         ST_Underline = 11
	ST_UnderlineDashDotHeavy    ST_Underline = 12
	ST_UnderlineDotDotDash      ST_Underline = 13
	ST_UnderlineDashDotDotHeavy ST_Underline = 14
	ST_UnderlineWave            ST_Underline = 15
	ST_UnderlineWavyHeavy       ST_Underline = 16
	ST_UnderlineWavyDouble      ST_Underline = 17
	ST_UnderlineNone            ST_Underline = 18
)

func (e ST_Underline) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_UnderlineUnset:
		attr.Value = ""
	case ST_UnderlineSingle:
		attr.Value = "single"
	case ST_UnderlineWords:
		attr.Value = "words"
	case ST_UnderlineDouble:
		attr.Value = "double"
	case ST_UnderlineThick:
		attr.Value = "thick"
	case ST_UnderlineDotted:
		attr.Value = "dotted"
	case ST_UnderlineDottedHeavy:
		attr.Value = "dottedHeavy"
	case ST_UnderlineDash:
		attr.Value = "dash"
	case ST_UnderlineDashedHeavy:
		attr.Value = "dashedHeavy"
	case ST_UnderlineDashLong:
		attr.Value = "dashLong"
	case ST_UnderlineDashLongHeavy:
		attr.Value = "dashLongHeavy"
	case ST_UnderlineDotDash:
		attr.Value = "dotDash"
	case ST_UnderlineDashDotHeavy:
		attr.Value = "dashDotHeavy"
	case ST_UnderlineDotDotDash:
		attr.Value = "dotDotDash"
	case ST_UnderlineDashDotDotHeavy:
		attr.Value = "dashDotDotHeavy"
	case ST_UnderlineWave:
		attr.Value = "wave"
	case ST_UnderlineWavyHeavy:
		attr.Value = "wavyHeavy"
	case ST_UnderlineWavyDouble:
		attr.Value = "wavyDouble"
	case ST_UnderlineNone:
		attr.Value = "none"
	}
	return attr, nil
}

func (e *ST_Underline) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "single":
		*e = 1
	case "words":
		*e = 2
	case "double":
		*e = 3
	case "thick":
		*e = 4
	case "dotted":
		*e = 5
	case "dottedHeavy":
		*e = 6
	case "dash":
		*e = 7
	case "dashedHeavy":
		*e = 8
	case "dashLong":
		*e = 9
	case "dashLongHeavy":
		*e = 10
	case "dotDash":
		*e = 11
	case "dashDotHeavy":
		*e = 12
	case "dotDotDash":
		*e = 13
	case "dashDotDotHeavy":
		*e = 14
	case "wave":
		*e = 15
	case "wavyHeavy":
		*e = 16
	case "wavyDouble":
		*e = 17
	case "none":
		*e = 18
	}
	return nil
}

func (m ST_Underline) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Underline) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "single":
			*m = 1
		case "words":
			*m = 2
		case "double":
			*m = 3
		case "thick":
			*m = 4
		case "dotted":
			*m = 5
		case "dottedHeavy":
			*m = 6
		case "dash":
			*m = 7
		case "dashedHeavy":
			*m = 8
		case "dashLong":
			*m = 9
		case "dashLongHeavy":
			*m = 10
		case "dotDash":
			*m = 11
		case "dashDotHeavy":
			*m = 12
		case "dotDotDash":
			*m = 13
		case "dashDotDotHeavy":
			*m = 14
		case "wave":
			*m = 15
		case "wavyHeavy":
			*m = 16
		case "wavyDouble":
			*m = 17
		case "none":
			*m = 18
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Underline) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "single"
	case 2:
		return "words"
	case 3:
		return "double"
	case 4:
		return "thick"
	case 5:
		return "dotted"
	case 6:
		return "dottedHeavy"
	case 7:
		return "dash"
	case 8:
		return "dashedHeavy"
	case 9:
		return "dashLong"
	case 10:
		return "dashLongHeavy"
	case 11:
		return "dotDash"
	case 12:
		return "dashDotHeavy"
	case 13:
		return "dotDotDash"
	case 14:
		return "dashDotDotHeavy"
	case 15:
		return "wave"
	case 16:
		return "wavyHeavy"
	case 17:
		return "wavyDouble"
	case 18:
		return "none"
	}
	return ""
}

func (m ST_Underline) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Underline) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextEffect byte

const (
	ST_TextEffectUnset           ST_TextEffect = 0
	ST_TextEffectBlinkBackground ST_TextEffect = 1
	ST_TextEffectLights          ST_TextEffect = 2
	ST_TextEffectAntsBlack       ST_TextEffect = 3
	ST_TextEffectAntsRed         ST_TextEffect = 4
	ST_TextEffectShimmer         ST_TextEffect = 5
	ST_TextEffectSparkle         ST_TextEffect = 6
	ST_TextEffectNone            ST_TextEffect = 7
)

func (e ST_TextEffect) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextEffectUnset:
		attr.Value = ""
	case ST_TextEffectBlinkBackground:
		attr.Value = "blinkBackground"
	case ST_TextEffectLights:
		attr.Value = "lights"
	case ST_TextEffectAntsBlack:
		attr.Value = "antsBlack"
	case ST_TextEffectAntsRed:
		attr.Value = "antsRed"
	case ST_TextEffectShimmer:
		attr.Value = "shimmer"
	case ST_TextEffectSparkle:
		attr.Value = "sparkle"
	case ST_TextEffectNone:
		attr.Value = "none"
	}
	return attr, nil
}

func (e *ST_TextEffect) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "blinkBackground":
		*e = 1
	case "lights":
		*e = 2
	case "antsBlack":
		*e = 3
	case "antsRed":
		*e = 4
	case "shimmer":
		*e = 5
	case "sparkle":
		*e = 6
	case "none":
		*e = 7
	}
	return nil
}

func (m ST_TextEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TextEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "blinkBackground":
			*m = 1
		case "lights":
			*m = 2
		case "antsBlack":
			*m = 3
		case "antsRed":
			*m = 4
		case "shimmer":
			*m = 5
		case "sparkle":
			*m = 6
		case "none":
			*m = 7
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_TextEffect) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "blinkBackground"
	case 2:
		return "lights"
	case 3:
		return "antsBlack"
	case 4:
		return "antsRed"
	case 5:
		return "shimmer"
	case 6:
		return "sparkle"
	case 7:
		return "none"
	}
	return ""
}

func (m ST_TextEffect) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TextEffect) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Border byte

const (
	ST_BorderUnset                  ST_Border = 0
	ST_BorderNil                    ST_Border = 1
	ST_BorderNone                   ST_Border = 2
	ST_BorderSingle                 ST_Border = 3
	ST_BorderThick                  ST_Border = 4
	ST_BorderDouble                 ST_Border = 5
	ST_BorderDotted                 ST_Border = 6
	ST_BorderDashed                 ST_Border = 7
	ST_BorderDotDash                ST_Border = 8
	ST_BorderDotDotDash             ST_Border = 9
	ST_BorderTriple                 ST_Border = 10
	ST_BorderThinThickSmallGap      ST_Border = 11
	ST_BorderThickThinSmallGap      ST_Border = 12
	ST_BorderThinThickThinSmallGap  ST_Border = 13
	ST_BorderThinThickMediumGap     ST_Border = 14
	ST_BorderThickThinMediumGap     ST_Border = 15
	ST_BorderThinThickThinMediumGap ST_Border = 16
	ST_BorderThinThickLargeGap      ST_Border = 17
	ST_BorderThickThinLargeGap      ST_Border = 18
	ST_BorderThinThickThinLargeGap  ST_Border = 19
	ST_BorderWave                   ST_Border = 20
	ST_BorderDoubleWave             ST_Border = 21
	ST_BorderDashSmallGap           ST_Border = 22
	ST_BorderDashDotStroked         ST_Border = 23
	ST_BorderThreeDEmboss           ST_Border = 24
	ST_BorderThreeDEngrave          ST_Border = 25
	ST_BorderOutset                 ST_Border = 26
	ST_BorderInset                  ST_Border = 27
	ST_BorderApples                 ST_Border = 28
	ST_BorderArchedScallops         ST_Border = 29
	ST_BorderBabyPacifier           ST_Border = 30
	ST_BorderBabyRattle             ST_Border = 31
	ST_BorderBalloons3Colors        ST_Border = 32
	ST_BorderBalloonsHotAir         ST_Border = 33
	ST_BorderBasicBlackDashes       ST_Border = 34
	ST_BorderBasicBlackDots         ST_Border = 35
	ST_BorderBasicBlackSquares      ST_Border = 36
	ST_BorderBasicThinLines         ST_Border = 37
	ST_BorderBasicWhiteDashes       ST_Border = 38
	ST_BorderBasicWhiteDots         ST_Border = 39
	ST_BorderBasicWhiteSquares      ST_Border = 40
	ST_BorderBasicWideInline        ST_Border = 41
	ST_BorderBasicWideMidline       ST_Border = 42
	ST_BorderBasicWideOutline       ST_Border = 43
	ST_BorderBats                   ST_Border = 44
	ST_BorderBirds                  ST_Border = 45
	ST_BorderBirdsFlight            ST_Border = 46
	ST_BorderCabins                 ST_Border = 47
	ST_BorderCakeSlice              ST_Border = 48
	ST_BorderCandyCorn              ST_Border = 49
	ST_BorderCelticKnotwork         ST_Border = 50
	ST_BorderCertificateBanner      ST_Border = 51
	ST_BorderChainLink              ST_Border = 52
	ST_BorderChampagneBottle        ST_Border = 53
	ST_BorderCheckedBarBlack        ST_Border = 54
	ST_BorderCheckedBarColor        ST_Border = 55
	ST_BorderCheckered              ST_Border = 56
	ST_BorderChristmasTree          ST_Border = 57
	ST_BorderCirclesLines           ST_Border = 58
	ST_BorderCirclesRectangles      ST_Border = 59
	ST_BorderClassicalWave          ST_Border = 60
	ST_BorderClocks                 ST_Border = 61
	ST_BorderCompass                ST_Border = 62
	ST_BorderConfetti               ST_Border = 63
	ST_BorderConfettiGrays          ST_Border = 64
	ST_BorderConfettiOutline        ST_Border = 65
	ST_BorderConfettiStreamers      ST_Border = 66
	ST_BorderConfettiWhite          ST_Border = 67
	ST_BorderCornerTriangles        ST_Border = 68
	ST_BorderCouponCutoutDashes     ST_Border = 69
	ST_BorderCouponCutoutDots       ST_Border = 70
	ST_BorderCrazyMaze              ST_Border = 71
	ST_BorderCreaturesButterfly     ST_Border = 72
	ST_BorderCreaturesFish          ST_Border = 73
	ST_BorderCreaturesInsects       ST_Border = 74
	ST_BorderCreaturesLadyBug       ST_Border = 75
	ST_BorderCrossStitch            ST_Border = 76
	ST_BorderCup                    ST_Border = 77
	ST_BorderDecoArch               ST_Border = 78
	ST_BorderDecoArchColor          ST_Border = 79
	ST_BorderDecoBlocks             ST_Border = 80
	ST_BorderDiamondsGray           ST_Border = 81
	ST_BorderDoubleD                ST_Border = 82
	ST_BorderDoubleDiamonds         ST_Border = 83
	ST_BorderEarth1                 ST_Border = 84
	ST_BorderEarth2                 ST_Border = 85
	ST_BorderEarth3                 ST_Border = 86
	ST_BorderEclipsingSquares1      ST_Border = 87
	ST_BorderEclipsingSquares2      ST_Border = 88
	ST_BorderEggsBlack              ST_Border = 89
	ST_BorderFans                   ST_Border = 90
	ST_BorderFilm                   ST_Border = 91
	ST_BorderFirecrackers           ST_Border = 92
	ST_BorderFlowersBlockPrint      ST_Border = 93
	ST_BorderFlowersDaisies         ST_Border = 94
	ST_BorderFlowersModern1         ST_Border = 95
	ST_BorderFlowersModern2         ST_Border = 96
	ST_BorderFlowersPansy           ST_Border = 97
	ST_BorderFlowersRedRose         ST_Border = 98
	ST_BorderFlowersRoses           ST_Border = 99
	ST_BorderFlowersTeacup          ST_Border = 100
	ST_BorderFlowersTiny            ST_Border = 101
	ST_BorderGems                   ST_Border = 102
	ST_BorderGingerbreadMan         ST_Border = 103
	ST_BorderGradient               ST_Border = 104
	ST_BorderHandmade1              ST_Border = 105
	ST_BorderHandmade2              ST_Border = 106
	ST_BorderHeartBalloon           ST_Border = 107
	ST_BorderHeartGray              ST_Border = 108
	ST_BorderHearts                 ST_Border = 109
	ST_BorderHeebieJeebies          ST_Border = 110
	ST_BorderHolly                  ST_Border = 111
	ST_BorderHouseFunky             ST_Border = 112
	ST_BorderHypnotic               ST_Border = 113
	ST_BorderIceCreamCones          ST_Border = 114
	ST_BorderLightBulb              ST_Border = 115
	ST_BorderLightning1             ST_Border = 116
	ST_BorderLightning2             ST_Border = 117
	ST_BorderMapPins                ST_Border = 118
	ST_BorderMapleLeaf              ST_Border = 119
	ST_BorderMapleMuffins           ST_Border = 120
	ST_BorderMarquee                ST_Border = 121
	ST_BorderMarqueeToothed         ST_Border = 122
	ST_BorderMoons                  ST_Border = 123
	ST_BorderMosaic                 ST_Border = 124
	ST_BorderMusicNotes             ST_Border = 125
	ST_BorderNorthwest              ST_Border = 126
	ST_BorderOvals                  ST_Border = 127
	ST_BorderPackages               ST_Border = 128
	ST_BorderPalmsBlack             ST_Border = 129
	ST_BorderPalmsColor             ST_Border = 130
	ST_BorderPaperClips             ST_Border = 131
	ST_BorderPapyrus                ST_Border = 132
	ST_BorderPartyFavor             ST_Border = 133
	ST_BorderPartyGlass             ST_Border = 134
	ST_BorderPencils                ST_Border = 135
	ST_BorderPeople                 ST_Border = 136
	ST_BorderPeopleWaving           ST_Border = 137
	ST_BorderPeopleHats             ST_Border = 138
	ST_BorderPoinsettias            ST_Border = 139
	ST_BorderPostageStamp           ST_Border = 140
	ST_BorderPumpkin1               ST_Border = 141
	ST_BorderPushPinNote2           ST_Border = 142
	ST_BorderPushPinNote1           ST_Border = 143
	ST_BorderPyramids               ST_Border = 144
	ST_BorderPyramidsAbove          ST_Border = 145
	ST_BorderQuadrants              ST_Border = 146
	ST_BorderRings                  ST_Border = 147
	ST_BorderSafari                 ST_Border = 148
	ST_BorderSawtooth               ST_Border = 149
	ST_BorderSawtoothGray           ST_Border = 150
	ST_BorderScaredCat              ST_Border = 151
	ST_BorderSeattle                ST_Border = 152
	ST_BorderShadowedSquares        ST_Border = 153
	ST_BorderSharksTeeth            ST_Border = 154
	ST_BorderShorebirdTracks        ST_Border = 155
	ST_BorderSkyrocket              ST_Border = 156
	ST_BorderSnowflakeFancy         ST_Border = 157
	ST_BorderSnowflakes             ST_Border = 158
	ST_BorderSombrero               ST_Border = 159
	ST_BorderSouthwest              ST_Border = 160
	ST_BorderStars                  ST_Border = 161
	ST_BorderStarsTop               ST_Border = 162
	ST_BorderStars3d                ST_Border = 163
	ST_BorderStarsBlack             ST_Border = 164
	ST_BorderStarsShadowed          ST_Border = 165
	ST_BorderSun                    ST_Border = 166
	ST_BorderSwirligig              ST_Border = 167
	ST_BorderTornPaper              ST_Border = 168
	ST_BorderTornPaperBlack         ST_Border = 169
	ST_BorderTrees                  ST_Border = 170
	ST_BorderTriangleParty          ST_Border = 171
	ST_BorderTriangles              ST_Border = 172
	ST_BorderTriangle1              ST_Border = 173
	ST_BorderTriangle2              ST_Border = 174
	ST_BorderTriangleCircle1        ST_Border = 175
	ST_BorderTriangleCircle2        ST_Border = 176
	ST_BorderShapes1                ST_Border = 177
	ST_BorderShapes2                ST_Border = 178
	ST_BorderTwistedLines1          ST_Border = 179
	ST_BorderTwistedLines2          ST_Border = 180
	ST_BorderVine                   ST_Border = 181
	ST_BorderWaveline               ST_Border = 182
	ST_BorderWeavingAngles          ST_Border = 183
	ST_BorderWeavingBraid           ST_Border = 184
	ST_BorderWeavingRibbon          ST_Border = 185
	ST_BorderWeavingStrips          ST_Border = 186
	ST_BorderWhiteFlowers           ST_Border = 187
	ST_BorderWoodwork               ST_Border = 188
	ST_BorderXIllusions             ST_Border = 189
	ST_BorderZanyTriangles          ST_Border = 190
	ST_BorderZigZag                 ST_Border = 191
	ST_BorderZigZagStitch           ST_Border = 192
	ST_BorderCustom                 ST_Border = 193
)

func (e ST_Border) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_BorderUnset:
		attr.Value = ""
	case ST_BorderNil:
		attr.Value = "nil"
	case ST_BorderNone:
		attr.Value = "none"
	case ST_BorderSingle:
		attr.Value = "single"
	case ST_BorderThick:
		attr.Value = "thick"
	case ST_BorderDouble:
		attr.Value = "double"
	case ST_BorderDotted:
		attr.Value = "dotted"
	case ST_BorderDashed:
		attr.Value = "dashed"
	case ST_BorderDotDash:
		attr.Value = "dotDash"
	case ST_BorderDotDotDash:
		attr.Value = "dotDotDash"
	case ST_BorderTriple:
		attr.Value = "triple"
	case ST_BorderThinThickSmallGap:
		attr.Value = "thinThickSmallGap"
	case ST_BorderThickThinSmallGap:
		attr.Value = "thickThinSmallGap"
	case ST_BorderThinThickThinSmallGap:
		attr.Value = "thinThickThinSmallGap"
	case ST_BorderThinThickMediumGap:
		attr.Value = "thinThickMediumGap"
	case ST_BorderThickThinMediumGap:
		attr.Value = "thickThinMediumGap"
	case ST_BorderThinThickThinMediumGap:
		attr.Value = "thinThickThinMediumGap"
	case ST_BorderThinThickLargeGap:
		attr.Value = "thinThickLargeGap"
	case ST_BorderThickThinLargeGap:
		attr.Value = "thickThinLargeGap"
	case ST_BorderThinThickThinLargeGap:
		attr.Value = "thinThickThinLargeGap"
	case ST_BorderWave:
		attr.Value = "wave"
	case ST_BorderDoubleWave:
		attr.Value = "doubleWave"
	case ST_BorderDashSmallGap:
		attr.Value = "dashSmallGap"
	case ST_BorderDashDotStroked:
		attr.Value = "dashDotStroked"
	case ST_BorderThreeDEmboss:
		attr.Value = "threeDEmboss"
	case ST_BorderThreeDEngrave:
		attr.Value = "threeDEngrave"
	case ST_BorderOutset:
		attr.Value = "outset"
	case ST_BorderInset:
		attr.Value = "inset"
	case ST_BorderApples:
		attr.Value = "apples"
	case ST_BorderArchedScallops:
		attr.Value = "archedScallops"
	case ST_BorderBabyPacifier:
		attr.Value = "babyPacifier"
	case ST_BorderBabyRattle:
		attr.Value = "babyRattle"
	case ST_BorderBalloons3Colors:
		attr.Value = "balloons3Colors"
	case ST_BorderBalloonsHotAir:
		attr.Value = "balloonsHotAir"
	case ST_BorderBasicBlackDashes:
		attr.Value = "basicBlackDashes"
	case ST_BorderBasicBlackDots:
		attr.Value = "basicBlackDots"
	case ST_BorderBasicBlackSquares:
		attr.Value = "basicBlackSquares"
	case ST_BorderBasicThinLines:
		attr.Value = "basicThinLines"
	case ST_BorderBasicWhiteDashes:
		attr.Value = "basicWhiteDashes"
	case ST_BorderBasicWhiteDots:
		attr.Value = "basicWhiteDots"
	case ST_BorderBasicWhiteSquares:
		attr.Value = "basicWhiteSquares"
	case ST_BorderBasicWideInline:
		attr.Value = "basicWideInline"
	case ST_BorderBasicWideMidline:
		attr.Value = "basicWideMidline"
	case ST_BorderBasicWideOutline:
		attr.Value = "basicWideOutline"
	case ST_BorderBats:
		attr.Value = "bats"
	case ST_BorderBirds:
		attr.Value = "birds"
	case ST_BorderBirdsFlight:
		attr.Value = "birdsFlight"
	case ST_BorderCabins:
		attr.Value = "cabins"
	case ST_BorderCakeSlice:
		attr.Value = "cakeSlice"
	case ST_BorderCandyCorn:
		attr.Value = "candyCorn"
	case ST_BorderCelticKnotwork:
		attr.Value = "celticKnotwork"
	case ST_BorderCertificateBanner:
		attr.Value = "certificateBanner"
	case ST_BorderChainLink:
		attr.Value = "chainLink"
	case ST_BorderChampagneBottle:
		attr.Value = "champagneBottle"
	case ST_BorderCheckedBarBlack:
		attr.Value = "checkedBarBlack"
	case ST_BorderCheckedBarColor:
		attr.Value = "checkedBarColor"
	case ST_BorderCheckered:
		attr.Value = "checkered"
	case ST_BorderChristmasTree:
		attr.Value = "christmasTree"
	case ST_BorderCirclesLines:
		attr.Value = "circlesLines"
	case ST_BorderCirclesRectangles:
		attr.Value = "circlesRectangles"
	case ST_BorderClassicalWave:
		attr.Value = "classicalWave"
	case ST_BorderClocks:
		attr.Value = "clocks"
	case ST_BorderCompass:
		attr.Value = "compass"
	case ST_BorderConfetti:
		attr.Value = "confetti"
	case ST_BorderConfettiGrays:
		attr.Value = "confettiGrays"
	case ST_BorderConfettiOutline:
		attr.Value = "confettiOutline"
	case ST_BorderConfettiStreamers:
		attr.Value = "confettiStreamers"
	case ST_BorderConfettiWhite:
		attr.Value = "confettiWhite"
	case ST_BorderCornerTriangles:
		attr.Value = "cornerTriangles"
	case ST_BorderCouponCutoutDashes:
		attr.Value = "couponCutoutDashes"
	case ST_BorderCouponCutoutDots:
		attr.Value = "couponCutoutDots"
	case ST_BorderCrazyMaze:
		attr.Value = "crazyMaze"
	case ST_BorderCreaturesButterfly:
		attr.Value = "creaturesButterfly"
	case ST_BorderCreaturesFish:
		attr.Value = "creaturesFish"
	case ST_BorderCreaturesInsects:
		attr.Value = "creaturesInsects"
	case ST_BorderCreaturesLadyBug:
		attr.Value = "creaturesLadyBug"
	case ST_BorderCrossStitch:
		attr.Value = "crossStitch"
	case ST_BorderCup:
		attr.Value = "cup"
	case ST_BorderDecoArch:
		attr.Value = "decoArch"
	case ST_BorderDecoArchColor:
		attr.Value = "decoArchColor"
	case ST_BorderDecoBlocks:
		attr.Value = "decoBlocks"
	case ST_BorderDiamondsGray:
		attr.Value = "diamondsGray"
	case ST_BorderDoubleD:
		attr.Value = "doubleD"
	case ST_BorderDoubleDiamonds:
		attr.Value = "doubleDiamonds"
	case ST_BorderEarth1:
		attr.Value = "earth1"
	case ST_BorderEarth2:
		attr.Value = "earth2"
	case ST_BorderEarth3:
		attr.Value = "earth3"
	case ST_BorderEclipsingSquares1:
		attr.Value = "eclipsingSquares1"
	case ST_BorderEclipsingSquares2:
		attr.Value = "eclipsingSquares2"
	case ST_BorderEggsBlack:
		attr.Value = "eggsBlack"
	case ST_BorderFans:
		attr.Value = "fans"
	case ST_BorderFilm:
		attr.Value = "film"
	case ST_BorderFirecrackers:
		attr.Value = "firecrackers"
	case ST_BorderFlowersBlockPrint:
		attr.Value = "flowersBlockPrint"
	case ST_BorderFlowersDaisies:
		attr.Value = "flowersDaisies"
	case ST_BorderFlowersModern1:
		attr.Value = "flowersModern1"
	case ST_BorderFlowersModern2:
		attr.Value = "flowersModern2"
	case ST_BorderFlowersPansy:
		attr.Value = "flowersPansy"
	case ST_BorderFlowersRedRose:
		attr.Value = "flowersRedRose"
	case ST_BorderFlowersRoses:
		attr.Value = "flowersRoses"
	case ST_BorderFlowersTeacup:
		attr.Value = "flowersTeacup"
	case ST_BorderFlowersTiny:
		attr.Value = "flowersTiny"
	case ST_BorderGems:
		attr.Value = "gems"
	case ST_BorderGingerbreadMan:
		attr.Value = "gingerbreadMan"
	case ST_BorderGradient:
		attr.Value = "gradient"
	case ST_BorderHandmade1:
		attr.Value = "handmade1"
	case ST_BorderHandmade2:
		attr.Value = "handmade2"
	case ST_BorderHeartBalloon:
		attr.Value = "heartBalloon"
	case ST_BorderHeartGray:
		attr.Value = "heartGray"
	case ST_BorderHearts:
		attr.Value = "hearts"
	case ST_BorderHeebieJeebies:
		attr.Value = "heebieJeebies"
	case ST_BorderHolly:
		attr.Value = "holly"
	case ST_BorderHouseFunky:
		attr.Value = "houseFunky"
	case ST_BorderHypnotic:
		attr.Value = "hypnotic"
	case ST_BorderIceCreamCones:
		attr.Value = "iceCreamCones"
	case ST_BorderLightBulb:
		attr.Value = "lightBulb"
	case ST_BorderLightning1:
		attr.Value = "lightning1"
	case ST_BorderLightning2:
		attr.Value = "lightning2"
	case ST_BorderMapPins:
		attr.Value = "mapPins"
	case ST_BorderMapleLeaf:
		attr.Value = "mapleLeaf"
	case ST_BorderMapleMuffins:
		attr.Value = "mapleMuffins"
	case ST_BorderMarquee:
		attr.Value = "marquee"
	case ST_BorderMarqueeToothed:
		attr.Value = "marqueeToothed"
	case ST_BorderMoons:
		attr.Value = "moons"
	case ST_BorderMosaic:
		attr.Value = "mosaic"
	case ST_BorderMusicNotes:
		attr.Value = "musicNotes"
	case ST_BorderNorthwest:
		attr.Value = "northwest"
	case ST_BorderOvals:
		attr.Value = "ovals"
	case ST_BorderPackages:
		attr.Value = "packages"
	case ST_BorderPalmsBlack:
		attr.Value = "palmsBlack"
	case ST_BorderPalmsColor:
		attr.Value = "palmsColor"
	case ST_BorderPaperClips:
		attr.Value = "paperClips"
	case ST_BorderPapyrus:
		attr.Value = "papyrus"
	case ST_BorderPartyFavor:
		attr.Value = "partyFavor"
	case ST_BorderPartyGlass:
		attr.Value = "partyGlass"
	case ST_BorderPencils:
		attr.Value = "pencils"
	case ST_BorderPeople:
		attr.Value = "people"
	case ST_BorderPeopleWaving:
		attr.Value = "peopleWaving"
	case ST_BorderPeopleHats:
		attr.Value = "peopleHats"
	case ST_BorderPoinsettias:
		attr.Value = "poinsettias"
	case ST_BorderPostageStamp:
		attr.Value = "postageStamp"
	case ST_BorderPumpkin1:
		attr.Value = "pumpkin1"
	case ST_BorderPushPinNote2:
		attr.Value = "pushPinNote2"
	case ST_BorderPushPinNote1:
		attr.Value = "pushPinNote1"
	case ST_BorderPyramids:
		attr.Value = "pyramids"
	case ST_BorderPyramidsAbove:
		attr.Value = "pyramidsAbove"
	case ST_BorderQuadrants:
		attr.Value = "quadrants"
	case ST_BorderRings:
		attr.Value = "rings"
	case ST_BorderSafari:
		attr.Value = "safari"
	case ST_BorderSawtooth:
		attr.Value = "sawtooth"
	case ST_BorderSawtoothGray:
		attr.Value = "sawtoothGray"
	case ST_BorderScaredCat:
		attr.Value = "scaredCat"
	case ST_BorderSeattle:
		attr.Value = "seattle"
	case ST_BorderShadowedSquares:
		attr.Value = "shadowedSquares"
	case ST_BorderSharksTeeth:
		attr.Value = "sharksTeeth"
	case ST_BorderShorebirdTracks:
		attr.Value = "shorebirdTracks"
	case ST_BorderSkyrocket:
		attr.Value = "skyrocket"
	case ST_BorderSnowflakeFancy:
		attr.Value = "snowflakeFancy"
	case ST_BorderSnowflakes:
		attr.Value = "snowflakes"
	case ST_BorderSombrero:
		attr.Value = "sombrero"
	case ST_BorderSouthwest:
		attr.Value = "southwest"
	case ST_BorderStars:
		attr.Value = "stars"
	case ST_BorderStarsTop:
		attr.Value = "starsTop"
	case ST_BorderStars3d:
		attr.Value = "stars3d"
	case ST_BorderStarsBlack:
		attr.Value = "starsBlack"
	case ST_BorderStarsShadowed:
		attr.Value = "starsShadowed"
	case ST_BorderSun:
		attr.Value = "sun"
	case ST_BorderSwirligig:
		attr.Value = "swirligig"
	case ST_BorderTornPaper:
		attr.Value = "tornPaper"
	case ST_BorderTornPaperBlack:
		attr.Value = "tornPaperBlack"
	case ST_BorderTrees:
		attr.Value = "trees"
	case ST_BorderTriangleParty:
		attr.Value = "triangleParty"
	case ST_BorderTriangles:
		attr.Value = "triangles"
	case ST_BorderTriangle1:
		attr.Value = "triangle1"
	case ST_BorderTriangle2:
		attr.Value = "triangle2"
	case ST_BorderTriangleCircle1:
		attr.Value = "triangleCircle1"
	case ST_BorderTriangleCircle2:
		attr.Value = "triangleCircle2"
	case ST_BorderShapes1:
		attr.Value = "shapes1"
	case ST_BorderShapes2:
		attr.Value = "shapes2"
	case ST_BorderTwistedLines1:
		attr.Value = "twistedLines1"
	case ST_BorderTwistedLines2:
		attr.Value = "twistedLines2"
	case ST_BorderVine:
		attr.Value = "vine"
	case ST_BorderWaveline:
		attr.Value = "waveline"
	case ST_BorderWeavingAngles:
		attr.Value = "weavingAngles"
	case ST_BorderWeavingBraid:
		attr.Value = "weavingBraid"
	case ST_BorderWeavingRibbon:
		attr.Value = "weavingRibbon"
	case ST_BorderWeavingStrips:
		attr.Value = "weavingStrips"
	case ST_BorderWhiteFlowers:
		attr.Value = "whiteFlowers"
	case ST_BorderWoodwork:
		attr.Value = "woodwork"
	case ST_BorderXIllusions:
		attr.Value = "xIllusions"
	case ST_BorderZanyTriangles:
		attr.Value = "zanyTriangles"
	case ST_BorderZigZag:
		attr.Value = "zigZag"
	case ST_BorderZigZagStitch:
		attr.Value = "zigZagStitch"
	case ST_BorderCustom:
		attr.Value = "custom"
	}
	return attr, nil
}

func (e *ST_Border) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "nil":
		*e = 1
	case "none":
		*e = 2
	case "single":
		*e = 3
	case "thick":
		*e = 4
	case "double":
		*e = 5
	case "dotted":
		*e = 6
	case "dashed":
		*e = 7
	case "dotDash":
		*e = 8
	case "dotDotDash":
		*e = 9
	case "triple":
		*e = 10
	case "thinThickSmallGap":
		*e = 11
	case "thickThinSmallGap":
		*e = 12
	case "thinThickThinSmallGap":
		*e = 13
	case "thinThickMediumGap":
		*e = 14
	case "thickThinMediumGap":
		*e = 15
	case "thinThickThinMediumGap":
		*e = 16
	case "thinThickLargeGap":
		*e = 17
	case "thickThinLargeGap":
		*e = 18
	case "thinThickThinLargeGap":
		*e = 19
	case "wave":
		*e = 20
	case "doubleWave":
		*e = 21
	case "dashSmallGap":
		*e = 22
	case "dashDotStroked":
		*e = 23
	case "threeDEmboss":
		*e = 24
	case "threeDEngrave":
		*e = 25
	case "outset":
		*e = 26
	case "inset":
		*e = 27
	case "apples":
		*e = 28
	case "archedScallops":
		*e = 29
	case "babyPacifier":
		*e = 30
	case "babyRattle":
		*e = 31
	case "balloons3Colors":
		*e = 32
	case "balloonsHotAir":
		*e = 33
	case "basicBlackDashes":
		*e = 34
	case "basicBlackDots":
		*e = 35
	case "basicBlackSquares":
		*e = 36
	case "basicThinLines":
		*e = 37
	case "basicWhiteDashes":
		*e = 38
	case "basicWhiteDots":
		*e = 39
	case "basicWhiteSquares":
		*e = 40
	case "basicWideInline":
		*e = 41
	case "basicWideMidline":
		*e = 42
	case "basicWideOutline":
		*e = 43
	case "bats":
		*e = 44
	case "birds":
		*e = 45
	case "birdsFlight":
		*e = 46
	case "cabins":
		*e = 47
	case "cakeSlice":
		*e = 48
	case "candyCorn":
		*e = 49
	case "celticKnotwork":
		*e = 50
	case "certificateBanner":
		*e = 51
	case "chainLink":
		*e = 52
	case "champagneBottle":
		*e = 53
	case "checkedBarBlack":
		*e = 54
	case "checkedBarColor":
		*e = 55
	case "checkered":
		*e = 56
	case "christmasTree":
		*e = 57
	case "circlesLines":
		*e = 58
	case "circlesRectangles":
		*e = 59
	case "classicalWave":
		*e = 60
	case "clocks":
		*e = 61
	case "compass":
		*e = 62
	case "confetti":
		*e = 63
	case "confettiGrays":
		*e = 64
	case "confettiOutline":
		*e = 65
	case "confettiStreamers":
		*e = 66
	case "confettiWhite":
		*e = 67
	case "cornerTriangles":
		*e = 68
	case "couponCutoutDashes":
		*e = 69
	case "couponCutoutDots":
		*e = 70
	case "crazyMaze":
		*e = 71
	case "creaturesButterfly":
		*e = 72
	case "creaturesFish":
		*e = 73
	case "creaturesInsects":
		*e = 74
	case "creaturesLadyBug":
		*e = 75
	case "crossStitch":
		*e = 76
	case "cup":
		*e = 77
	case "decoArch":
		*e = 78
	case "decoArchColor":
		*e = 79
	case "decoBlocks":
		*e = 80
	case "diamondsGray":
		*e = 81
	case "doubleD":
		*e = 82
	case "doubleDiamonds":
		*e = 83
	case "earth1":
		*e = 84
	case "earth2":
		*e = 85
	case "earth3":
		*e = 86
	case "eclipsingSquares1":
		*e = 87
	case "eclipsingSquares2":
		*e = 88
	case "eggsBlack":
		*e = 89
	case "fans":
		*e = 90
	case "film":
		*e = 91
	case "firecrackers":
		*e = 92
	case "flowersBlockPrint":
		*e = 93
	case "flowersDaisies":
		*e = 94
	case "flowersModern1":
		*e = 95
	case "flowersModern2":
		*e = 96
	case "flowersPansy":
		*e = 97
	case "flowersRedRose":
		*e = 98
	case "flowersRoses":
		*e = 99
	case "flowersTeacup":
		*e = 100
	case "flowersTiny":
		*e = 101
	case "gems":
		*e = 102
	case "gingerbreadMan":
		*e = 103
	case "gradient":
		*e = 104
	case "handmade1":
		*e = 105
	case "handmade2":
		*e = 106
	case "heartBalloon":
		*e = 107
	case "heartGray":
		*e = 108
	case "hearts":
		*e = 109
	case "heebieJeebies":
		*e = 110
	case "holly":
		*e = 111
	case "houseFunky":
		*e = 112
	case "hypnotic":
		*e = 113
	case "iceCreamCones":
		*e = 114
	case "lightBulb":
		*e = 115
	case "lightning1":
		*e = 116
	case "lightning2":
		*e = 117
	case "mapPins":
		*e = 118
	case "mapleLeaf":
		*e = 119
	case "mapleMuffins":
		*e = 120
	case "marquee":
		*e = 121
	case "marqueeToothed":
		*e = 122
	case "moons":
		*e = 123
	case "mosaic":
		*e = 124
	case "musicNotes":
		*e = 125
	case "northwest":
		*e = 126
	case "ovals":
		*e = 127
	case "packages":
		*e = 128
	case "palmsBlack":
		*e = 129
	case "palmsColor":
		*e = 130
	case "paperClips":
		*e = 131
	case "papyrus":
		*e = 132
	case "partyFavor":
		*e = 133
	case "partyGlass":
		*e = 134
	case "pencils":
		*e = 135
	case "people":
		*e = 136
	case "peopleWaving":
		*e = 137
	case "peopleHats":
		*e = 138
	case "poinsettias":
		*e = 139
	case "postageStamp":
		*e = 140
	case "pumpkin1":
		*e = 141
	case "pushPinNote2":
		*e = 142
	case "pushPinNote1":
		*e = 143
	case "pyramids":
		*e = 144
	case "pyramidsAbove":
		*e = 145
	case "quadrants":
		*e = 146
	case "rings":
		*e = 147
	case "safari":
		*e = 148
	case "sawtooth":
		*e = 149
	case "sawtoothGray":
		*e = 150
	case "scaredCat":
		*e = 151
	case "seattle":
		*e = 152
	case "shadowedSquares":
		*e = 153
	case "sharksTeeth":
		*e = 154
	case "shorebirdTracks":
		*e = 155
	case "skyrocket":
		*e = 156
	case "snowflakeFancy":
		*e = 157
	case "snowflakes":
		*e = 158
	case "sombrero":
		*e = 159
	case "southwest":
		*e = 160
	case "stars":
		*e = 161
	case "starsTop":
		*e = 162
	case "stars3d":
		*e = 163
	case "starsBlack":
		*e = 164
	case "starsShadowed":
		*e = 165
	case "sun":
		*e = 166
	case "swirligig":
		*e = 167
	case "tornPaper":
		*e = 168
	case "tornPaperBlack":
		*e = 169
	case "trees":
		*e = 170
	case "triangleParty":
		*e = 171
	case "triangles":
		*e = 172
	case "triangle1":
		*e = 173
	case "triangle2":
		*e = 174
	case "triangleCircle1":
		*e = 175
	case "triangleCircle2":
		*e = 176
	case "shapes1":
		*e = 177
	case "shapes2":
		*e = 178
	case "twistedLines1":
		*e = 179
	case "twistedLines2":
		*e = 180
	case "vine":
		*e = 181
	case "waveline":
		*e = 182
	case "weavingAngles":
		*e = 183
	case "weavingBraid":
		*e = 184
	case "weavingRibbon":
		*e = 185
	case "weavingStrips":
		*e = 186
	case "whiteFlowers":
		*e = 187
	case "woodwork":
		*e = 188
	case "xIllusions":
		*e = 189
	case "zanyTriangles":
		*e = 190
	case "zigZag":
		*e = 191
	case "zigZagStitch":
		*e = 192
	case "custom":
		*e = 193
	}
	return nil
}

func (m ST_Border) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Border) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "nil":
			*m = 1
		case "none":
			*m = 2
		case "single":
			*m = 3
		case "thick":
			*m = 4
		case "double":
			*m = 5
		case "dotted":
			*m = 6
		case "dashed":
			*m = 7
		case "dotDash":
			*m = 8
		case "dotDotDash":
			*m = 9
		case "triple":
			*m = 10
		case "thinThickSmallGap":
			*m = 11
		case "thickThinSmallGap":
			*m = 12
		case "thinThickThinSmallGap":
			*m = 13
		case "thinThickMediumGap":
			*m = 14
		case "thickThinMediumGap":
			*m = 15
		case "thinThickThinMediumGap":
			*m = 16
		case "thinThickLargeGap":
			*m = 17
		case "thickThinLargeGap":
			*m = 18
		case "thinThickThinLargeGap":
			*m = 19
		case "wave":
			*m = 20
		case "doubleWave":
			*m = 21
		case "dashSmallGap":
			*m = 22
		case "dashDotStroked":
			*m = 23
		case "threeDEmboss":
			*m = 24
		case "threeDEngrave":
			*m = 25
		case "outset":
			*m = 26
		case "inset":
			*m = 27
		case "apples":
			*m = 28
		case "archedScallops":
			*m = 29
		case "babyPacifier":
			*m = 30
		case "babyRattle":
			*m = 31
		case "balloons3Colors":
			*m = 32
		case "balloonsHotAir":
			*m = 33
		case "basicBlackDashes":
			*m = 34
		case "basicBlackDots":
			*m = 35
		case "basicBlackSquares":
			*m = 36
		case "basicThinLines":
			*m = 37
		case "basicWhiteDashes":
			*m = 38
		case "basicWhiteDots":
			*m = 39
		case "basicWhiteSquares":
			*m = 40
		case "basicWideInline":
			*m = 41
		case "basicWideMidline":
			*m = 42
		case "basicWideOutline":
			*m = 43
		case "bats":
			*m = 44
		case "birds":
			*m = 45
		case "birdsFlight":
			*m = 46
		case "cabins":
			*m = 47
		case "cakeSlice":
			*m = 48
		case "candyCorn":
			*m = 49
		case "celticKnotwork":
			*m = 50
		case "certificateBanner":
			*m = 51
		case "chainLink":
			*m = 52
		case "champagneBottle":
			*m = 53
		case "checkedBarBlack":
			*m = 54
		case "checkedBarColor":
			*m = 55
		case "checkered":
			*m = 56
		case "christmasTree":
			*m = 57
		case "circlesLines":
			*m = 58
		case "circlesRectangles":
			*m = 59
		case "classicalWave":
			*m = 60
		case "clocks":
			*m = 61
		case "compass":
			*m = 62
		case "confetti":
			*m = 63
		case "confettiGrays":
			*m = 64
		case "confettiOutline":
			*m = 65
		case "confettiStreamers":
			*m = 66
		case "confettiWhite":
			*m = 67
		case "cornerTriangles":
			*m = 68
		case "couponCutoutDashes":
			*m = 69
		case "couponCutoutDots":
			*m = 70
		case "crazyMaze":
			*m = 71
		case "creaturesButterfly":
			*m = 72
		case "creaturesFish":
			*m = 73
		case "creaturesInsects":
			*m = 74
		case "creaturesLadyBug":
			*m = 75
		case "crossStitch":
			*m = 76
		case "cup":
			*m = 77
		case "decoArch":
			*m = 78
		case "decoArchColor":
			*m = 79
		case "decoBlocks":
			*m = 80
		case "diamondsGray":
			*m = 81
		case "doubleD":
			*m = 82
		case "doubleDiamonds":
			*m = 83
		case "earth1":
			*m = 84
		case "earth2":
			*m = 85
		case "earth3":
			*m = 86
		case "eclipsingSquares1":
			*m = 87
		case "eclipsingSquares2":
			*m = 88
		case "eggsBlack":
			*m = 89
		case "fans":
			*m = 90
		case "film":
			*m = 91
		case "firecrackers":
			*m = 92
		case "flowersBlockPrint":
			*m = 93
		case "flowersDaisies":
			*m = 94
		case "flowersModern1":
			*m = 95
		case "flowersModern2":
			*m = 96
		case "flowersPansy":
			*m = 97
		case "flowersRedRose":
			*m = 98
		case "flowersRoses":
			*m = 99
		case "flowersTeacup":
			*m = 100
		case "flowersTiny":
			*m = 101
		case "gems":
			*m = 102
		case "gingerbreadMan":
			*m = 103
		case "gradient":
			*m = 104
		case "handmade1":
			*m = 105
		case "handmade2":
			*m = 106
		case "heartBalloon":
			*m = 107
		case "heartGray":
			*m = 108
		case "hearts":
			*m = 109
		case "heebieJeebies":
			*m = 110
		case "holly":
			*m = 111
		case "houseFunky":
			*m = 112
		case "hypnotic":
			*m = 113
		case "iceCreamCones":
			*m = 114
		case "lightBulb":
			*m = 115
		case "lightning1":
			*m = 116
		case "lightning2":
			*m = 117
		case "mapPins":
			*m = 118
		case "mapleLeaf":
			*m = 119
		case "mapleMuffins":
			*m = 120
		case "marquee":
			*m = 121
		case "marqueeToothed":
			*m = 122
		case "moons":
			*m = 123
		case "mosaic":
			*m = 124
		case "musicNotes":
			*m = 125
		case "northwest":
			*m = 126
		case "ovals":
			*m = 127
		case "packages":
			*m = 128
		case "palmsBlack":
			*m = 129
		case "palmsColor":
			*m = 130
		case "paperClips":
			*m = 131
		case "papyrus":
			*m = 132
		case "partyFavor":
			*m = 133
		case "partyGlass":
			*m = 134
		case "pencils":
			*m = 135
		case "people":
			*m = 136
		case "peopleWaving":
			*m = 137
		case "peopleHats":
			*m = 138
		case "poinsettias":
			*m = 139
		case "postageStamp":
			*m = 140
		case "pumpkin1":
			*m = 141
		case "pushPinNote2":
			*m = 142
		case "pushPinNote1":
			*m = 143
		case "pyramids":
			*m = 144
		case "pyramidsAbove":
			*m = 145
		case "quadrants":
			*m = 146
		case "rings":
			*m = 147
		case "safari":
			*m = 148
		case "sawtooth":
			*m = 149
		case "sawtoothGray":
			*m = 150
		case "scaredCat":
			*m = 151
		case "seattle":
			*m = 152
		case "shadowedSquares":
			*m = 153
		case "sharksTeeth":
			*m = 154
		case "shorebirdTracks":
			*m = 155
		case "skyrocket":
			*m = 156
		case "snowflakeFancy":
			*m = 157
		case "snowflakes":
			*m = 158
		case "sombrero":
			*m = 159
		case "southwest":
			*m = 160
		case "stars":
			*m = 161
		case "starsTop":
			*m = 162
		case "stars3d":
			*m = 163
		case "starsBlack":
			*m = 164
		case "starsShadowed":
			*m = 165
		case "sun":
			*m = 166
		case "swirligig":
			*m = 167
		case "tornPaper":
			*m = 168
		case "tornPaperBlack":
			*m = 169
		case "trees":
			*m = 170
		case "triangleParty":
			*m = 171
		case "triangles":
			*m = 172
		case "triangle1":
			*m = 173
		case "triangle2":
			*m = 174
		case "triangleCircle1":
			*m = 175
		case "triangleCircle2":
			*m = 176
		case "shapes1":
			*m = 177
		case "shapes2":
			*m = 178
		case "twistedLines1":
			*m = 179
		case "twistedLines2":
			*m = 180
		case "vine":
			*m = 181
		case "waveline":
			*m = 182
		case "weavingAngles":
			*m = 183
		case "weavingBraid":
			*m = 184
		case "weavingRibbon":
			*m = 185
		case "weavingStrips":
			*m = 186
		case "whiteFlowers":
			*m = 187
		case "woodwork":
			*m = 188
		case "xIllusions":
			*m = 189
		case "zanyTriangles":
			*m = 190
		case "zigZag":
			*m = 191
		case "zigZagStitch":
			*m = 192
		case "custom":
			*m = 193
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Border) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "nil"
	case 2:
		return "none"
	case 3:
		return "single"
	case 4:
		return "thick"
	case 5:
		return "double"
	case 6:
		return "dotted"
	case 7:
		return "dashed"
	case 8:
		return "dotDash"
	case 9:
		return "dotDotDash"
	case 10:
		return "triple"
	case 11:
		return "thinThickSmallGap"
	case 12:
		return "thickThinSmallGap"
	case 13:
		return "thinThickThinSmallGap"
	case 14:
		return "thinThickMediumGap"
	case 15:
		return "thickThinMediumGap"
	case 16:
		return "thinThickThinMediumGap"
	case 17:
		return "thinThickLargeGap"
	case 18:
		return "thickThinLargeGap"
	case 19:
		return "thinThickThinLargeGap"
	case 20:
		return "wave"
	case 21:
		return "doubleWave"
	case 22:
		return "dashSmallGap"
	case 23:
		return "dashDotStroked"
	case 24:
		return "threeDEmboss"
	case 25:
		return "threeDEngrave"
	case 26:
		return "outset"
	case 27:
		return "inset"
	case 28:
		return "apples"
	case 29:
		return "archedScallops"
	case 30:
		return "babyPacifier"
	case 31:
		return "babyRattle"
	case 32:
		return "balloons3Colors"
	case 33:
		return "balloonsHotAir"
	case 34:
		return "basicBlackDashes"
	case 35:
		return "basicBlackDots"
	case 36:
		return "basicBlackSquares"
	case 37:
		return "basicThinLines"
	case 38:
		return "basicWhiteDashes"
	case 39:
		return "basicWhiteDots"
	case 40:
		return "basicWhiteSquares"
	case 41:
		return "basicWideInline"
	case 42:
		return "basicWideMidline"
	case 43:
		return "basicWideOutline"
	case 44:
		return "bats"
	case 45:
		return "birds"
	case 46:
		return "birdsFlight"
	case 47:
		return "cabins"
	case 48:
		return "cakeSlice"
	case 49:
		return "candyCorn"
	case 50:
		return "celticKnotwork"
	case 51:
		return "certificateBanner"
	case 52:
		return "chainLink"
	case 53:
		return "champagneBottle"
	case 54:
		return "checkedBarBlack"
	case 55:
		return "checkedBarColor"
	case 56:
		return "checkered"
	case 57:
		return "christmasTree"
	case 58:
		return "circlesLines"
	case 59:
		return "circlesRectangles"
	case 60:
		return "classicalWave"
	case 61:
		return "clocks"
	case 62:
		return "compass"
	case 63:
		return "confetti"
	case 64:
		return "confettiGrays"
	case 65:
		return "confettiOutline"
	case 66:
		return "confettiStreamers"
	case 67:
		return "confettiWhite"
	case 68:
		return "cornerTriangles"
	case 69:
		return "couponCutoutDashes"
	case 70:
		return "couponCutoutDots"
	case 71:
		return "crazyMaze"
	case 72:
		return "creaturesButterfly"
	case 73:
		return "creaturesFish"
	case 74:
		return "creaturesInsects"
	case 75:
		return "creaturesLadyBug"
	case 76:
		return "crossStitch"
	case 77:
		return "cup"
	case 78:
		return "decoArch"
	case 79:
		return "decoArchColor"
	case 80:
		return "decoBlocks"
	case 81:
		return "diamondsGray"
	case 82:
		return "doubleD"
	case 83:
		return "doubleDiamonds"
	case 84:
		return "earth1"
	case 85:
		return "earth2"
	case 86:
		return "earth3"
	case 87:
		return "eclipsingSquares1"
	case 88:
		return "eclipsingSquares2"
	case 89:
		return "eggsBlack"
	case 90:
		return "fans"
	case 91:
		return "film"
	case 92:
		return "firecrackers"
	case 93:
		return "flowersBlockPrint"
	case 94:
		return "flowersDaisies"
	case 95:
		return "flowersModern1"
	case 96:
		return "flowersModern2"
	case 97:
		return "flowersPansy"
	case 98:
		return "flowersRedRose"
	case 99:
		return "flowersRoses"
	case 100:
		return "flowersTeacup"
	case 101:
		return "flowersTiny"
	case 102:
		return "gems"
	case 103:
		return "gingerbreadMan"
	case 104:
		return "gradient"
	case 105:
		return "handmade1"
	case 106:
		return "handmade2"
	case 107:
		return "heartBalloon"
	case 108:
		return "heartGray"
	case 109:
		return "hearts"
	case 110:
		return "heebieJeebies"
	case 111:
		return "holly"
	case 112:
		return "houseFunky"
	case 113:
		return "hypnotic"
	case 114:
		return "iceCreamCones"
	case 115:
		return "lightBulb"
	case 116:
		return "lightning1"
	case 117:
		return "lightning2"
	case 118:
		return "mapPins"
	case 119:
		return "mapleLeaf"
	case 120:
		return "mapleMuffins"
	case 121:
		return "marquee"
	case 122:
		return "marqueeToothed"
	case 123:
		return "moons"
	case 124:
		return "mosaic"
	case 125:
		return "musicNotes"
	case 126:
		return "northwest"
	case 127:
		return "ovals"
	case 128:
		return "packages"
	case 129:
		return "palmsBlack"
	case 130:
		return "palmsColor"
	case 131:
		return "paperClips"
	case 132:
		return "papyrus"
	case 133:
		return "partyFavor"
	case 134:
		return "partyGlass"
	case 135:
		return "pencils"
	case 136:
		return "people"
	case 137:
		return "peopleWaving"
	case 138:
		return "peopleHats"
	case 139:
		return "poinsettias"
	case 140:
		return "postageStamp"
	case 141:
		return "pumpkin1"
	case 142:
		return "pushPinNote2"
	case 143:
		return "pushPinNote1"
	case 144:
		return "pyramids"
	case 145:
		return "pyramidsAbove"
	case 146:
		return "quadrants"
	case 147:
		return "rings"
	case 148:
		return "safari"
	case 149:
		return "sawtooth"
	case 150:
		return "sawtoothGray"
	case 151:
		return "scaredCat"
	case 152:
		return "seattle"
	case 153:
		return "shadowedSquares"
	case 154:
		return "sharksTeeth"
	case 155:
		return "shorebirdTracks"
	case 156:
		return "skyrocket"
	case 157:
		return "snowflakeFancy"
	case 158:
		return "snowflakes"
	case 159:
		return "sombrero"
	case 160:
		return "southwest"
	case 161:
		return "stars"
	case 162:
		return "starsTop"
	case 163:
		return "stars3d"
	case 164:
		return "starsBlack"
	case 165:
		return "starsShadowed"
	case 166:
		return "sun"
	case 167:
		return "swirligig"
	case 168:
		return "tornPaper"
	case 169:
		return "tornPaperBlack"
	case 170:
		return "trees"
	case 171:
		return "triangleParty"
	case 172:
		return "triangles"
	case 173:
		return "triangle1"
	case 174:
		return "triangle2"
	case 175:
		return "triangleCircle1"
	case 176:
		return "triangleCircle2"
	case 177:
		return "shapes1"
	case 178:
		return "shapes2"
	case 179:
		return "twistedLines1"
	case 180:
		return "twistedLines2"
	case 181:
		return "vine"
	case 182:
		return "waveline"
	case 183:
		return "weavingAngles"
	case 184:
		return "weavingBraid"
	case 185:
		return "weavingRibbon"
	case 186:
		return "weavingStrips"
	case 187:
		return "whiteFlowers"
	case 188:
		return "woodwork"
	case 189:
		return "xIllusions"
	case 190:
		return "zanyTriangles"
	case 191:
		return "zigZag"
	case 192:
		return "zigZagStitch"
	case 193:
		return "custom"
	}
	return ""
}

func (m ST_Border) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Border) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 192, 193:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Shd byte

const (
	ST_ShdUnset                 ST_Shd = 0
	ST_ShdNil                   ST_Shd = 1
	ST_ShdClear                 ST_Shd = 2
	ST_ShdSolid                 ST_Shd = 3
	ST_ShdHorzStripe            ST_Shd = 4
	ST_ShdVertStripe            ST_Shd = 5
	ST_ShdReverseDiagStripe     ST_Shd = 6
	ST_ShdDiagStripe            ST_Shd = 7
	ST_ShdHorzCross             ST_Shd = 8
	ST_ShdDiagCross             ST_Shd = 9
	ST_ShdThinHorzStripe        ST_Shd = 10
	ST_ShdThinVertStripe        ST_Shd = 11
	ST_ShdThinReverseDiagStripe ST_Shd = 12
	ST_ShdThinDiagStripe        ST_Shd = 13
	ST_ShdThinHorzCross         ST_Shd = 14
	ST_ShdThinDiagCross         ST_Shd = 15
	ST_ShdPct5                  ST_Shd = 16
	ST_ShdPct10                 ST_Shd = 17
	ST_ShdPct12                 ST_Shd = 18
	ST_ShdPct15                 ST_Shd = 19
	ST_ShdPct20                 ST_Shd = 20
	ST_ShdPct25                 ST_Shd = 21
	ST_ShdPct30                 ST_Shd = 22
	ST_ShdPct35                 ST_Shd = 23
	ST_ShdPct37                 ST_Shd = 24
	ST_ShdPct40                 ST_Shd = 25
	ST_ShdPct45                 ST_Shd = 26
	ST_ShdPct50                 ST_Shd = 27
	ST_ShdPct55                 ST_Shd = 28
	ST_ShdPct60                 ST_Shd = 29
	ST_ShdPct62                 ST_Shd = 30
	ST_ShdPct65                 ST_Shd = 31
	ST_ShdPct70                 ST_Shd = 32
	ST_ShdPct75                 ST_Shd = 33
	ST_ShdPct80                 ST_Shd = 34
	ST_ShdPct85                 ST_Shd = 35
	ST_ShdPct87                 ST_Shd = 36
	ST_ShdPct90                 ST_Shd = 37
	ST_ShdPct95                 ST_Shd = 38
)

func (e ST_Shd) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ShdUnset:
		attr.Value = ""
	case ST_ShdNil:
		attr.Value = "nil"
	case ST_ShdClear:
		attr.Value = "clear"
	case ST_ShdSolid:
		attr.Value = "solid"
	case ST_ShdHorzStripe:
		attr.Value = "horzStripe"
	case ST_ShdVertStripe:
		attr.Value = "vertStripe"
	case ST_ShdReverseDiagStripe:
		attr.Value = "reverseDiagStripe"
	case ST_ShdDiagStripe:
		attr.Value = "diagStripe"
	case ST_ShdHorzCross:
		attr.Value = "horzCross"
	case ST_ShdDiagCross:
		attr.Value = "diagCross"
	case ST_ShdThinHorzStripe:
		attr.Value = "thinHorzStripe"
	case ST_ShdThinVertStripe:
		attr.Value = "thinVertStripe"
	case ST_ShdThinReverseDiagStripe:
		attr.Value = "thinReverseDiagStripe"
	case ST_ShdThinDiagStripe:
		attr.Value = "thinDiagStripe"
	case ST_ShdThinHorzCross:
		attr.Value = "thinHorzCross"
	case ST_ShdThinDiagCross:
		attr.Value = "thinDiagCross"
	case ST_ShdPct5:
		attr.Value = "pct5"
	case ST_ShdPct10:
		attr.Value = "pct10"
	case ST_ShdPct12:
		attr.Value = "pct12"
	case ST_ShdPct15:
		attr.Value = "pct15"
	case ST_ShdPct20:
		attr.Value = "pct20"
	case ST_ShdPct25:
		attr.Value = "pct25"
	case ST_ShdPct30:
		attr.Value = "pct30"
	case ST_ShdPct35:
		attr.Value = "pct35"
	case ST_ShdPct37:
		attr.Value = "pct37"
	case ST_ShdPct40:
		attr.Value = "pct40"
	case ST_ShdPct45:
		attr.Value = "pct45"
	case ST_ShdPct50:
		attr.Value = "pct50"
	case ST_ShdPct55:
		attr.Value = "pct55"
	case ST_ShdPct60:
		attr.Value = "pct60"
	case ST_ShdPct62:
		attr.Value = "pct62"
	case ST_ShdPct65:
		attr.Value = "pct65"
	case ST_ShdPct70:
		attr.Value = "pct70"
	case ST_ShdPct75:
		attr.Value = "pct75"
	case ST_ShdPct80:
		attr.Value = "pct80"
	case ST_ShdPct85:
		attr.Value = "pct85"
	case ST_ShdPct87:
		attr.Value = "pct87"
	case ST_ShdPct90:
		attr.Value = "pct90"
	case ST_ShdPct95:
		attr.Value = "pct95"
	}
	return attr, nil
}

func (e *ST_Shd) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "nil":
		*e = 1
	case "clear":
		*e = 2
	case "solid":
		*e = 3
	case "horzStripe":
		*e = 4
	case "vertStripe":
		*e = 5
	case "reverseDiagStripe":
		*e = 6
	case "diagStripe":
		*e = 7
	case "horzCross":
		*e = 8
	case "diagCross":
		*e = 9
	case "thinHorzStripe":
		*e = 10
	case "thinVertStripe":
		*e = 11
	case "thinReverseDiagStripe":
		*e = 12
	case "thinDiagStripe":
		*e = 13
	case "thinHorzCross":
		*e = 14
	case "thinDiagCross":
		*e = 15
	case "pct5":
		*e = 16
	case "pct10":
		*e = 17
	case "pct12":
		*e = 18
	case "pct15":
		*e = 19
	case "pct20":
		*e = 20
	case "pct25":
		*e = 21
	case "pct30":
		*e = 22
	case "pct35":
		*e = 23
	case "pct37":
		*e = 24
	case "pct40":
		*e = 25
	case "pct45":
		*e = 26
	case "pct50":
		*e = 27
	case "pct55":
		*e = 28
	case "pct60":
		*e = 29
	case "pct62":
		*e = 30
	case "pct65":
		*e = 31
	case "pct70":
		*e = 32
	case "pct75":
		*e = 33
	case "pct80":
		*e = 34
	case "pct85":
		*e = 35
	case "pct87":
		*e = 36
	case "pct90":
		*e = 37
	case "pct95":
		*e = 38
	}
	return nil
}

func (m ST_Shd) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Shd) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "nil":
			*m = 1
		case "clear":
			*m = 2
		case "solid":
			*m = 3
		case "horzStripe":
			*m = 4
		case "vertStripe":
			*m = 5
		case "reverseDiagStripe":
			*m = 6
		case "diagStripe":
			*m = 7
		case "horzCross":
			*m = 8
		case "diagCross":
			*m = 9
		case "thinHorzStripe":
			*m = 10
		case "thinVertStripe":
			*m = 11
		case "thinReverseDiagStripe":
			*m = 12
		case "thinDiagStripe":
			*m = 13
		case "thinHorzCross":
			*m = 14
		case "thinDiagCross":
			*m = 15
		case "pct5":
			*m = 16
		case "pct10":
			*m = 17
		case "pct12":
			*m = 18
		case "pct15":
			*m = 19
		case "pct20":
			*m = 20
		case "pct25":
			*m = 21
		case "pct30":
			*m = 22
		case "pct35":
			*m = 23
		case "pct37":
			*m = 24
		case "pct40":
			*m = 25
		case "pct45":
			*m = 26
		case "pct50":
			*m = 27
		case "pct55":
			*m = 28
		case "pct60":
			*m = 29
		case "pct62":
			*m = 30
		case "pct65":
			*m = 31
		case "pct70":
			*m = 32
		case "pct75":
			*m = 33
		case "pct80":
			*m = 34
		case "pct85":
			*m = 35
		case "pct87":
			*m = 36
		case "pct90":
			*m = 37
		case "pct95":
			*m = 38
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Shd) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "nil"
	case 2:
		return "clear"
	case 3:
		return "solid"
	case 4:
		return "horzStripe"
	case 5:
		return "vertStripe"
	case 6:
		return "reverseDiagStripe"
	case 7:
		return "diagStripe"
	case 8:
		return "horzCross"
	case 9:
		return "diagCross"
	case 10:
		return "thinHorzStripe"
	case 11:
		return "thinVertStripe"
	case 12:
		return "thinReverseDiagStripe"
	case 13:
		return "thinDiagStripe"
	case 14:
		return "thinHorzCross"
	case 15:
		return "thinDiagCross"
	case 16:
		return "pct5"
	case 17:
		return "pct10"
	case 18:
		return "pct12"
	case 19:
		return "pct15"
	case 20:
		return "pct20"
	case 21:
		return "pct25"
	case 22:
		return "pct30"
	case 23:
		return "pct35"
	case 24:
		return "pct37"
	case 25:
		return "pct40"
	case 26:
		return "pct45"
	case 27:
		return "pct50"
	case 28:
		return "pct55"
	case 29:
		return "pct60"
	case 30:
		return "pct62"
	case 31:
		return "pct65"
	case 32:
		return "pct70"
	case 33:
		return "pct75"
	case 34:
		return "pct80"
	case 35:
		return "pct85"
	case 36:
		return "pct87"
	case 37:
		return "pct90"
	case 38:
		return "pct95"
	}
	return ""
}

func (m ST_Shd) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Shd) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Em byte

const (
	ST_EmUnset    ST_Em = 0
	ST_EmNone     ST_Em = 1
	ST_EmDot      ST_Em = 2
	ST_EmComma    ST_Em = 3
	ST_EmCircle   ST_Em = 4
	ST_EmUnderDot ST_Em = 5
)

func (e ST_Em) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_EmUnset:
		attr.Value = ""
	case ST_EmNone:
		attr.Value = "none"
	case ST_EmDot:
		attr.Value = "dot"
	case ST_EmComma:
		attr.Value = "comma"
	case ST_EmCircle:
		attr.Value = "circle"
	case ST_EmUnderDot:
		attr.Value = "underDot"
	}
	return attr, nil
}

func (e *ST_Em) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "dot":
		*e = 2
	case "comma":
		*e = 3
	case "circle":
		*e = 4
	case "underDot":
		*e = 5
	}
	return nil
}

func (m ST_Em) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Em) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "dot":
			*m = 2
		case "comma":
			*m = 3
		case "circle":
			*m = 4
		case "underDot":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Em) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "dot"
	case 3:
		return "comma"
	case 4:
		return "circle"
	case 5:
		return "underDot"
	}
	return ""
}

func (m ST_Em) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Em) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_CombineBrackets byte

const (
	ST_CombineBracketsUnset  ST_CombineBrackets = 0
	ST_CombineBracketsNone   ST_CombineBrackets = 1
	ST_CombineBracketsRound  ST_CombineBrackets = 2
	ST_CombineBracketsSquare ST_CombineBrackets = 3
	ST_CombineBracketsAngle  ST_CombineBrackets = 4
	ST_CombineBracketsCurly  ST_CombineBrackets = 5
)

func (e ST_CombineBrackets) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_CombineBracketsUnset:
		attr.Value = ""
	case ST_CombineBracketsNone:
		attr.Value = "none"
	case ST_CombineBracketsRound:
		attr.Value = "round"
	case ST_CombineBracketsSquare:
		attr.Value = "square"
	case ST_CombineBracketsAngle:
		attr.Value = "angle"
	case ST_CombineBracketsCurly:
		attr.Value = "curly"
	}
	return attr, nil
}

func (e *ST_CombineBrackets) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "round":
		*e = 2
	case "square":
		*e = 3
	case "angle":
		*e = 4
	case "curly":
		*e = 5
	}
	return nil
}

func (m ST_CombineBrackets) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_CombineBrackets) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "round":
			*m = 2
		case "square":
			*m = 3
		case "angle":
			*m = 4
		case "curly":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_CombineBrackets) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "round"
	case 3:
		return "square"
	case 4:
		return "angle"
	case 5:
		return "curly"
	}
	return ""
}

func (m ST_CombineBrackets) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_CombineBrackets) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_HeightRule byte

const (
	ST_HeightRuleUnset   ST_HeightRule = 0
	ST_HeightRuleAuto    ST_HeightRule = 1
	ST_HeightRuleExact   ST_HeightRule = 2
	ST_HeightRuleAtLeast ST_HeightRule = 3
)

func (e ST_HeightRule) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_HeightRuleUnset:
		attr.Value = ""
	case ST_HeightRuleAuto:
		attr.Value = "auto"
	case ST_HeightRuleExact:
		attr.Value = "exact"
	case ST_HeightRuleAtLeast:
		attr.Value = "atLeast"
	}
	return attr, nil
}

func (e *ST_HeightRule) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "auto":
		*e = 1
	case "exact":
		*e = 2
	case "atLeast":
		*e = 3
	}
	return nil
}

func (m ST_HeightRule) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_HeightRule) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "auto":
			*m = 1
		case "exact":
			*m = 2
		case "atLeast":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_HeightRule) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "auto"
	case 2:
		return "exact"
	case 3:
		return "atLeast"
	}
	return ""
}

func (m ST_HeightRule) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_HeightRule) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Wrap byte

const (
	ST_WrapUnset     ST_Wrap = 0
	ST_WrapAuto      ST_Wrap = 1
	ST_WrapNotBeside ST_Wrap = 2
	ST_WrapAround    ST_Wrap = 3
	ST_WrapTight     ST_Wrap = 4
	ST_WrapThrough   ST_Wrap = 5
	ST_WrapNone      ST_Wrap = 6
)

func (e ST_Wrap) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_WrapUnset:
		attr.Value = ""
	case ST_WrapAuto:
		attr.Value = "auto"
	case ST_WrapNotBeside:
		attr.Value = "notBeside"
	case ST_WrapAround:
		attr.Value = "around"
	case ST_WrapTight:
		attr.Value = "tight"
	case ST_WrapThrough:
		attr.Value = "through"
	case ST_WrapNone:
		attr.Value = "none"
	}
	return attr, nil
}

func (e *ST_Wrap) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "auto":
		*e = 1
	case "notBeside":
		*e = 2
	case "around":
		*e = 3
	case "tight":
		*e = 4
	case "through":
		*e = 5
	case "none":
		*e = 6
	}
	return nil
}

func (m ST_Wrap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Wrap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "auto":
			*m = 1
		case "notBeside":
			*m = 2
		case "around":
			*m = 3
		case "tight":
			*m = 4
		case "through":
			*m = 5
		case "none":
			*m = 6
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Wrap) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "auto"
	case 2:
		return "notBeside"
	case 3:
		return "around"
	case 4:
		return "tight"
	case 5:
		return "through"
	case 6:
		return "none"
	}
	return ""
}

func (m ST_Wrap) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Wrap) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_VAnchor byte

const (
	ST_VAnchorUnset  ST_VAnchor = 0
	ST_VAnchorText   ST_VAnchor = 1
	ST_VAnchorMargin ST_VAnchor = 2
	ST_VAnchorPage   ST_VAnchor = 3
)

func (e ST_VAnchor) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_VAnchorUnset:
		attr.Value = ""
	case ST_VAnchorText:
		attr.Value = "text"
	case ST_VAnchorMargin:
		attr.Value = "margin"
	case ST_VAnchorPage:
		attr.Value = "page"
	}
	return attr, nil
}

func (e *ST_VAnchor) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "text":
		*e = 1
	case "margin":
		*e = 2
	case "page":
		*e = 3
	}
	return nil
}

func (m ST_VAnchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_VAnchor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "text":
			*m = 1
		case "margin":
			*m = 2
		case "page":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_VAnchor) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "text"
	case 2:
		return "margin"
	case 3:
		return "page"
	}
	return ""
}

func (m ST_VAnchor) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_VAnchor) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_HAnchor byte

const (
	ST_HAnchorUnset  ST_HAnchor = 0
	ST_HAnchorText   ST_HAnchor = 1
	ST_HAnchorMargin ST_HAnchor = 2
	ST_HAnchorPage   ST_HAnchor = 3
)

func (e ST_HAnchor) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_HAnchorUnset:
		attr.Value = ""
	case ST_HAnchorText:
		attr.Value = "text"
	case ST_HAnchorMargin:
		attr.Value = "margin"
	case ST_HAnchorPage:
		attr.Value = "page"
	}
	return attr, nil
}

func (e *ST_HAnchor) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "text":
		*e = 1
	case "margin":
		*e = 2
	case "page":
		*e = 3
	}
	return nil
}

func (m ST_HAnchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_HAnchor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "text":
			*m = 1
		case "margin":
			*m = 2
		case "page":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_HAnchor) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "text"
	case 2:
		return "margin"
	case 3:
		return "page"
	}
	return ""
}

func (m ST_HAnchor) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_HAnchor) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DropCap byte

const (
	ST_DropCapUnset  ST_DropCap = 0
	ST_DropCapNone   ST_DropCap = 1
	ST_DropCapDrop   ST_DropCap = 2
	ST_DropCapMargin ST_DropCap = 3
)

func (e ST_DropCap) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DropCapUnset:
		attr.Value = ""
	case ST_DropCapNone:
		attr.Value = "none"
	case ST_DropCapDrop:
		attr.Value = "drop"
	case ST_DropCapMargin:
		attr.Value = "margin"
	}
	return attr, nil
}

func (e *ST_DropCap) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "drop":
		*e = 2
	case "margin":
		*e = 3
	}
	return nil
}

func (m ST_DropCap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_DropCap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "drop":
			*m = 2
		case "margin":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_DropCap) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "drop"
	case 3:
		return "margin"
	}
	return ""
}

func (m ST_DropCap) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_DropCap) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TabJc byte

const (
	ST_TabJcUnset   ST_TabJc = 0
	ST_TabJcClear   ST_TabJc = 1
	ST_TabJcStart   ST_TabJc = 2
	ST_TabJcCenter  ST_TabJc = 3
	ST_TabJcEnd     ST_TabJc = 4
	ST_TabJcDecimal ST_TabJc = 5
	ST_TabJcBar     ST_TabJc = 6
	ST_TabJcNum     ST_TabJc = 7
	ST_TabJcLeft    ST_TabJc = 8
	ST_TabJcRight   ST_TabJc = 9
)

func (e ST_TabJc) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TabJcUnset:
		attr.Value = ""
	case ST_TabJcClear:
		attr.Value = "clear"
	case ST_TabJcStart:
		attr.Value = "start"
	case ST_TabJcCenter:
		attr.Value = "center"
	case ST_TabJcEnd:
		attr.Value = "end"
	case ST_TabJcDecimal:
		attr.Value = "decimal"
	case ST_TabJcBar:
		attr.Value = "bar"
	case ST_TabJcNum:
		attr.Value = "num"
	case ST_TabJcLeft:
		attr.Value = "left"
	case ST_TabJcRight:
		attr.Value = "right"
	}
	return attr, nil
}

func (e *ST_TabJc) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "clear":
		*e = 1
	case "start":
		*e = 2
	case "center":
		*e = 3
	case "end":
		*e = 4
	case "decimal":
		*e = 5
	case "bar":
		*e = 6
	case "num":
		*e = 7
	case "left":
		*e = 8
	case "right":
		*e = 9
	}
	return nil
}

func (m ST_TabJc) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TabJc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "clear":
			*m = 1
		case "start":
			*m = 2
		case "center":
			*m = 3
		case "end":
			*m = 4
		case "decimal":
			*m = 5
		case "bar":
			*m = 6
		case "num":
			*m = 7
		case "left":
			*m = 8
		case "right":
			*m = 9
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_TabJc) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "clear"
	case 2:
		return "start"
	case 3:
		return "center"
	case 4:
		return "end"
	case 5:
		return "decimal"
	case 6:
		return "bar"
	case 7:
		return "num"
	case 8:
		return "left"
	case 9:
		return "right"
	}
	return ""
}

func (m ST_TabJc) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TabJc) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TabTlc byte

const (
	ST_TabTlcUnset      ST_TabTlc = 0
	ST_TabTlcNone       ST_TabTlc = 1
	ST_TabTlcDot        ST_TabTlc = 2
	ST_TabTlcHyphen     ST_TabTlc = 3
	ST_TabTlcUnderscore ST_TabTlc = 4
	ST_TabTlcHeavy      ST_TabTlc = 5
	ST_TabTlcMiddleDot  ST_TabTlc = 6
)

func (e ST_TabTlc) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TabTlcUnset:
		attr.Value = ""
	case ST_TabTlcNone:
		attr.Value = "none"
	case ST_TabTlcDot:
		attr.Value = "dot"
	case ST_TabTlcHyphen:
		attr.Value = "hyphen"
	case ST_TabTlcUnderscore:
		attr.Value = "underscore"
	case ST_TabTlcHeavy:
		attr.Value = "heavy"
	case ST_TabTlcMiddleDot:
		attr.Value = "middleDot"
	}
	return attr, nil
}

func (e *ST_TabTlc) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "dot":
		*e = 2
	case "hyphen":
		*e = 3
	case "underscore":
		*e = 4
	case "heavy":
		*e = 5
	case "middleDot":
		*e = 6
	}
	return nil
}

func (m ST_TabTlc) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TabTlc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "dot":
			*m = 2
		case "hyphen":
			*m = 3
		case "underscore":
			*m = 4
		case "heavy":
			*m = 5
		case "middleDot":
			*m = 6
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_TabTlc) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "dot"
	case 3:
		return "hyphen"
	case 4:
		return "underscore"
	case 5:
		return "heavy"
	case 6:
		return "middleDot"
	}
	return ""
}

func (m ST_TabTlc) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TabTlc) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_LineSpacingRule byte

const (
	ST_LineSpacingRuleUnset   ST_LineSpacingRule = 0
	ST_LineSpacingRuleAuto    ST_LineSpacingRule = 1
	ST_LineSpacingRuleExact   ST_LineSpacingRule = 2
	ST_LineSpacingRuleAtLeast ST_LineSpacingRule = 3
)

func (e ST_LineSpacingRule) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_LineSpacingRuleUnset:
		attr.Value = ""
	case ST_LineSpacingRuleAuto:
		attr.Value = "auto"
	case ST_LineSpacingRuleExact:
		attr.Value = "exact"
	case ST_LineSpacingRuleAtLeast:
		attr.Value = "atLeast"
	}
	return attr, nil
}

func (e *ST_LineSpacingRule) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "auto":
		*e = 1
	case "exact":
		*e = 2
	case "atLeast":
		*e = 3
	}
	return nil
}

func (m ST_LineSpacingRule) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_LineSpacingRule) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "auto":
			*m = 1
		case "exact":
			*m = 2
		case "atLeast":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_LineSpacingRule) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "auto"
	case 2:
		return "exact"
	case 3:
		return "atLeast"
	}
	return ""
}

func (m ST_LineSpacingRule) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_LineSpacingRule) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Jc byte

const (
	ST_JcUnset          ST_Jc = 0
	ST_JcStart          ST_Jc = 1
	ST_JcCenter         ST_Jc = 2
	ST_JcEnd            ST_Jc = 3
	ST_JcBoth           ST_Jc = 4
	ST_JcMediumKashida  ST_Jc = 5
	ST_JcDistribute     ST_Jc = 6
	ST_JcNumTab         ST_Jc = 7
	ST_JcHighKashida    ST_Jc = 8
	ST_JcLowKashida     ST_Jc = 9
	ST_JcThaiDistribute ST_Jc = 10
	ST_JcLeft           ST_Jc = 11
	ST_JcRight          ST_Jc = 12
)

func (e ST_Jc) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_JcUnset:
		attr.Value = ""
	case ST_JcStart:
		attr.Value = "start"
	case ST_JcCenter:
		attr.Value = "center"
	case ST_JcEnd:
		attr.Value = "end"
	case ST_JcBoth:
		attr.Value = "both"
	case ST_JcMediumKashida:
		attr.Value = "mediumKashida"
	case ST_JcDistribute:
		attr.Value = "distribute"
	case ST_JcNumTab:
		attr.Value = "numTab"
	case ST_JcHighKashida:
		attr.Value = "highKashida"
	case ST_JcLowKashida:
		attr.Value = "lowKashida"
	case ST_JcThaiDistribute:
		attr.Value = "thaiDistribute"
	case ST_JcLeft:
		attr.Value = "left"
	case ST_JcRight:
		attr.Value = "right"
	}
	return attr, nil
}

func (e *ST_Jc) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "start":
		*e = 1
	case "center":
		*e = 2
	case "end":
		*e = 3
	case "both":
		*e = 4
	case "mediumKashida":
		*e = 5
	case "distribute":
		*e = 6
	case "numTab":
		*e = 7
	case "highKashida":
		*e = 8
	case "lowKashida":
		*e = 9
	case "thaiDistribute":
		*e = 10
	case "left":
		*e = 11
	case "right":
		*e = 12
	}
	return nil
}

func (m ST_Jc) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Jc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "start":
			*m = 1
		case "center":
			*m = 2
		case "end":
			*m = 3
		case "both":
			*m = 4
		case "mediumKashida":
			*m = 5
		case "distribute":
			*m = 6
		case "numTab":
			*m = 7
		case "highKashida":
			*m = 8
		case "lowKashida":
			*m = 9
		case "thaiDistribute":
			*m = 10
		case "left":
			*m = 11
		case "right":
			*m = 12
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Jc) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "start"
	case 2:
		return "center"
	case 3:
		return "end"
	case 4:
		return "both"
	case 5:
		return "mediumKashida"
	case 6:
		return "distribute"
	case 7:
		return "numTab"
	case 8:
		return "highKashida"
	case 9:
		return "lowKashida"
	case 10:
		return "thaiDistribute"
	case 11:
		return "left"
	case 12:
		return "right"
	}
	return ""
}

func (m ST_Jc) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Jc) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_JcTable byte

const (
	ST_JcTableUnset  ST_JcTable = 0
	ST_JcTableCenter ST_JcTable = 1
	ST_JcTableEnd    ST_JcTable = 2
	ST_JcTableLeft   ST_JcTable = 3
	ST_JcTableRight  ST_JcTable = 4
	ST_JcTableStart  ST_JcTable = 5
)

func (e ST_JcTable) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_JcTableUnset:
		attr.Value = ""
	case ST_JcTableCenter:
		attr.Value = "center"
	case ST_JcTableEnd:
		attr.Value = "end"
	case ST_JcTableLeft:
		attr.Value = "left"
	case ST_JcTableRight:
		attr.Value = "right"
	case ST_JcTableStart:
		attr.Value = "start"
	}
	return attr, nil
}

func (e *ST_JcTable) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "center":
		*e = 1
	case "end":
		*e = 2
	case "left":
		*e = 3
	case "right":
		*e = 4
	case "start":
		*e = 5
	}
	return nil
}

func (m ST_JcTable) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_JcTable) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "center":
			*m = 1
		case "end":
			*m = 2
		case "left":
			*m = 3
		case "right":
			*m = 4
		case "start":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_JcTable) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "center"
	case 2:
		return "end"
	case 3:
		return "left"
	case 4:
		return "right"
	case 5:
		return "start"
	}
	return ""
}

func (m ST_JcTable) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_JcTable) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_View byte

const (
	ST_ViewUnset       ST_View = 0
	ST_ViewNone        ST_View = 1
	ST_ViewPrint       ST_View = 2
	ST_ViewOutline     ST_View = 3
	ST_ViewMasterPages ST_View = 4
	ST_ViewNormal      ST_View = 5
	ST_ViewWeb         ST_View = 6
)

func (e ST_View) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ViewUnset:
		attr.Value = ""
	case ST_ViewNone:
		attr.Value = "none"
	case ST_ViewPrint:
		attr.Value = "print"
	case ST_ViewOutline:
		attr.Value = "outline"
	case ST_ViewMasterPages:
		attr.Value = "masterPages"
	case ST_ViewNormal:
		attr.Value = "normal"
	case ST_ViewWeb:
		attr.Value = "web"
	}
	return attr, nil
}

func (e *ST_View) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "print":
		*e = 2
	case "outline":
		*e = 3
	case "masterPages":
		*e = 4
	case "normal":
		*e = 5
	case "web":
		*e = 6
	}
	return nil
}

func (m ST_View) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_View) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "print":
			*m = 2
		case "outline":
			*m = 3
		case "masterPages":
			*m = 4
		case "normal":
			*m = 5
		case "web":
			*m = 6
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_View) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "print"
	case 3:
		return "outline"
	case 4:
		return "masterPages"
	case 5:
		return "normal"
	case 6:
		return "web"
	}
	return ""
}

func (m ST_View) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_View) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Zoom byte

const (
	ST_ZoomUnset    ST_Zoom = 0
	ST_ZoomNone     ST_Zoom = 1
	ST_ZoomFullPage ST_Zoom = 2
	ST_ZoomBestFit  ST_Zoom = 3
	ST_ZoomTextFit  ST_Zoom = 4
)

func (e ST_Zoom) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ZoomUnset:
		attr.Value = ""
	case ST_ZoomNone:
		attr.Value = "none"
	case ST_ZoomFullPage:
		attr.Value = "fullPage"
	case ST_ZoomBestFit:
		attr.Value = "bestFit"
	case ST_ZoomTextFit:
		attr.Value = "textFit"
	}
	return attr, nil
}

func (e *ST_Zoom) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "fullPage":
		*e = 2
	case "bestFit":
		*e = 3
	case "textFit":
		*e = 4
	}
	return nil
}

func (m ST_Zoom) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Zoom) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "fullPage":
			*m = 2
		case "bestFit":
			*m = 3
		case "textFit":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Zoom) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "fullPage"
	case 3:
		return "bestFit"
	case 4:
		return "textFit"
	}
	return ""
}

func (m ST_Zoom) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Zoom) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Proof byte

const (
	ST_ProofUnset ST_Proof = 0
	ST_ProofClean ST_Proof = 1
	ST_ProofDirty ST_Proof = 2
)

func (e ST_Proof) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ProofUnset:
		attr.Value = ""
	case ST_ProofClean:
		attr.Value = "clean"
	case ST_ProofDirty:
		attr.Value = "dirty"
	}
	return attr, nil
}

func (e *ST_Proof) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "clean":
		*e = 1
	case "dirty":
		*e = 2
	}
	return nil
}

func (m ST_Proof) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Proof) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "clean":
			*m = 1
		case "dirty":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Proof) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "clean"
	case 2:
		return "dirty"
	}
	return ""
}

func (m ST_Proof) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Proof) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DocProtect byte

const (
	ST_DocProtectUnset          ST_DocProtect = 0
	ST_DocProtectNone           ST_DocProtect = 1
	ST_DocProtectReadOnly       ST_DocProtect = 2
	ST_DocProtectComments       ST_DocProtect = 3
	ST_DocProtectTrackedChanges ST_DocProtect = 4
	ST_DocProtectForms          ST_DocProtect = 5
)

func (e ST_DocProtect) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DocProtectUnset:
		attr.Value = ""
	case ST_DocProtectNone:
		attr.Value = "none"
	case ST_DocProtectReadOnly:
		attr.Value = "readOnly"
	case ST_DocProtectComments:
		attr.Value = "comments"
	case ST_DocProtectTrackedChanges:
		attr.Value = "trackedChanges"
	case ST_DocProtectForms:
		attr.Value = "forms"
	}
	return attr, nil
}

func (e *ST_DocProtect) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "readOnly":
		*e = 2
	case "comments":
		*e = 3
	case "trackedChanges":
		*e = 4
	case "forms":
		*e = 5
	}
	return nil
}

func (m ST_DocProtect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_DocProtect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "readOnly":
			*m = 2
		case "comments":
			*m = 3
		case "trackedChanges":
			*m = 4
		case "forms":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_DocProtect) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "readOnly"
	case 3:
		return "comments"
	case 4:
		return "trackedChanges"
	case 5:
		return "forms"
	}
	return ""
}

func (m ST_DocProtect) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_DocProtect) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_MailMergeDocType byte

const (
	ST_MailMergeDocTypeUnset         ST_MailMergeDocType = 0
	ST_MailMergeDocTypeCatalog       ST_MailMergeDocType = 1
	ST_MailMergeDocTypeEnvelopes     ST_MailMergeDocType = 2
	ST_MailMergeDocTypeMailingLabels ST_MailMergeDocType = 3
	ST_MailMergeDocTypeFormLetters   ST_MailMergeDocType = 4
	ST_MailMergeDocTypeEmail         ST_MailMergeDocType = 5
	ST_MailMergeDocTypeFax           ST_MailMergeDocType = 6
)

func (e ST_MailMergeDocType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_MailMergeDocTypeUnset:
		attr.Value = ""
	case ST_MailMergeDocTypeCatalog:
		attr.Value = "catalog"
	case ST_MailMergeDocTypeEnvelopes:
		attr.Value = "envelopes"
	case ST_MailMergeDocTypeMailingLabels:
		attr.Value = "mailingLabels"
	case ST_MailMergeDocTypeFormLetters:
		attr.Value = "formLetters"
	case ST_MailMergeDocTypeEmail:
		attr.Value = "email"
	case ST_MailMergeDocTypeFax:
		attr.Value = "fax"
	}
	return attr, nil
}

func (e *ST_MailMergeDocType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "catalog":
		*e = 1
	case "envelopes":
		*e = 2
	case "mailingLabels":
		*e = 3
	case "formLetters":
		*e = 4
	case "email":
		*e = 5
	case "fax":
		*e = 6
	}
	return nil
}

func (m ST_MailMergeDocType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_MailMergeDocType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "catalog":
			*m = 1
		case "envelopes":
			*m = 2
		case "mailingLabels":
			*m = 3
		case "formLetters":
			*m = 4
		case "email":
			*m = 5
		case "fax":
			*m = 6
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_MailMergeDocType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "catalog"
	case 2:
		return "envelopes"
	case 3:
		return "mailingLabels"
	case 4:
		return "formLetters"
	case 5:
		return "email"
	case 6:
		return "fax"
	}
	return ""
}

func (m ST_MailMergeDocType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_MailMergeDocType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_MailMergeDest byte

const (
	ST_MailMergeDestUnset       ST_MailMergeDest = 0
	ST_MailMergeDestNewDocument ST_MailMergeDest = 1
	ST_MailMergeDestPrinter     ST_MailMergeDest = 2
	ST_MailMergeDestEmail       ST_MailMergeDest = 3
	ST_MailMergeDestFax         ST_MailMergeDest = 4
)

func (e ST_MailMergeDest) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_MailMergeDestUnset:
		attr.Value = ""
	case ST_MailMergeDestNewDocument:
		attr.Value = "newDocument"
	case ST_MailMergeDestPrinter:
		attr.Value = "printer"
	case ST_MailMergeDestEmail:
		attr.Value = "email"
	case ST_MailMergeDestFax:
		attr.Value = "fax"
	}
	return attr, nil
}

func (e *ST_MailMergeDest) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "newDocument":
		*e = 1
	case "printer":
		*e = 2
	case "email":
		*e = 3
	case "fax":
		*e = 4
	}
	return nil
}

func (m ST_MailMergeDest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_MailMergeDest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "newDocument":
			*m = 1
		case "printer":
			*m = 2
		case "email":
			*m = 3
		case "fax":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_MailMergeDest) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "newDocument"
	case 2:
		return "printer"
	case 3:
		return "email"
	case 4:
		return "fax"
	}
	return ""
}

func (m ST_MailMergeDest) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_MailMergeDest) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_MailMergeOdsoFMDFieldType byte

const (
	ST_MailMergeOdsoFMDFieldTypeUnset    ST_MailMergeOdsoFMDFieldType = 0
	ST_MailMergeOdsoFMDFieldTypeNull     ST_MailMergeOdsoFMDFieldType = 1
	ST_MailMergeOdsoFMDFieldTypeDbColumn ST_MailMergeOdsoFMDFieldType = 2
)

func (e ST_MailMergeOdsoFMDFieldType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_MailMergeOdsoFMDFieldTypeUnset:
		attr.Value = ""
	case ST_MailMergeOdsoFMDFieldTypeNull:
		attr.Value = "null"
	case ST_MailMergeOdsoFMDFieldTypeDbColumn:
		attr.Value = "dbColumn"
	}
	return attr, nil
}

func (e *ST_MailMergeOdsoFMDFieldType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "null":
		*e = 1
	case "dbColumn":
		*e = 2
	}
	return nil
}

func (m ST_MailMergeOdsoFMDFieldType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_MailMergeOdsoFMDFieldType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "null":
			*m = 1
		case "dbColumn":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_MailMergeOdsoFMDFieldType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "null"
	case 2:
		return "dbColumn"
	}
	return ""
}

func (m ST_MailMergeOdsoFMDFieldType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_MailMergeOdsoFMDFieldType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextDirection byte

const (
	ST_TextDirectionUnset ST_TextDirection = 0
	ST_TextDirectionTb    ST_TextDirection = 1
	ST_TextDirectionRl    ST_TextDirection = 2
	ST_TextDirectionLr    ST_TextDirection = 3
	ST_TextDirectionTbV   ST_TextDirection = 4
	ST_TextDirectionRlV   ST_TextDirection = 5
	ST_TextDirectionLrV   ST_TextDirection = 6
	ST_TextDirectionBtLr  ST_TextDirection = 7
	ST_TextDirectionLrTb  ST_TextDirection = 8
	ST_TextDirectionLrTbV ST_TextDirection = 9
	ST_TextDirectionTbLrV ST_TextDirection = 10
	ST_TextDirectionTbRl  ST_TextDirection = 11
	ST_TextDirectionTbRlV ST_TextDirection = 12
)

func (e ST_TextDirection) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextDirectionUnset:
		attr.Value = ""
	case ST_TextDirectionTb:
		attr.Value = "tb"
	case ST_TextDirectionRl:
		attr.Value = "rl"
	case ST_TextDirectionLr:
		attr.Value = "lr"
	case ST_TextDirectionTbV:
		attr.Value = "tbV"
	case ST_TextDirectionRlV:
		attr.Value = "rlV"
	case ST_TextDirectionLrV:
		attr.Value = "lrV"
	case ST_TextDirectionBtLr:
		attr.Value = "btLr"
	case ST_TextDirectionLrTb:
		attr.Value = "lrTb"
	case ST_TextDirectionLrTbV:
		attr.Value = "lrTbV"
	case ST_TextDirectionTbLrV:
		attr.Value = "tbLrV"
	case ST_TextDirectionTbRl:
		attr.Value = "tbRl"
	case ST_TextDirectionTbRlV:
		attr.Value = "tbRlV"
	}
	return attr, nil
}

func (e *ST_TextDirection) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "tb":
		*e = 1
	case "rl":
		*e = 2
	case "lr":
		*e = 3
	case "tbV":
		*e = 4
	case "rlV":
		*e = 5
	case "lrV":
		*e = 6
	case "btLr":
		*e = 7
	case "lrTb":
		*e = 8
	case "lrTbV":
		*e = 9
	case "tbLrV":
		*e = 10
	case "tbRl":
		*e = 11
	case "tbRlV":
		*e = 12
	}
	return nil
}

func (m ST_TextDirection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TextDirection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "tb":
			*m = 1
		case "rl":
			*m = 2
		case "lr":
			*m = 3
		case "tbV":
			*m = 4
		case "rlV":
			*m = 5
		case "lrV":
			*m = 6
		case "btLr":
			*m = 7
		case "lrTb":
			*m = 8
		case "lrTbV":
			*m = 9
		case "tbLrV":
			*m = 10
		case "tbRl":
			*m = 11
		case "tbRlV":
			*m = 12
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_TextDirection) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "tb"
	case 2:
		return "rl"
	case 3:
		return "lr"
	case 4:
		return "tbV"
	case 5:
		return "rlV"
	case 6:
		return "lrV"
	case 7:
		return "btLr"
	case 8:
		return "lrTb"
	case 9:
		return "lrTbV"
	case 10:
		return "tbLrV"
	case 11:
		return "tbRl"
	case 12:
		return "tbRlV"
	}
	return ""
}

func (m ST_TextDirection) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TextDirection) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextAlignment byte

const (
	ST_TextAlignmentUnset    ST_TextAlignment = 0
	ST_TextAlignmentTop      ST_TextAlignment = 1
	ST_TextAlignmentCenter   ST_TextAlignment = 2
	ST_TextAlignmentBaseline ST_TextAlignment = 3
	ST_TextAlignmentBottom   ST_TextAlignment = 4
	ST_TextAlignmentAuto     ST_TextAlignment = 5
)

func (e ST_TextAlignment) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextAlignmentUnset:
		attr.Value = ""
	case ST_TextAlignmentTop:
		attr.Value = "top"
	case ST_TextAlignmentCenter:
		attr.Value = "center"
	case ST_TextAlignmentBaseline:
		attr.Value = "baseline"
	case ST_TextAlignmentBottom:
		attr.Value = "bottom"
	case ST_TextAlignmentAuto:
		attr.Value = "auto"
	}
	return attr, nil
}

func (e *ST_TextAlignment) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "top":
		*e = 1
	case "center":
		*e = 2
	case "baseline":
		*e = 3
	case "bottom":
		*e = 4
	case "auto":
		*e = 5
	}
	return nil
}

func (m ST_TextAlignment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TextAlignment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "top":
			*m = 1
		case "center":
			*m = 2
		case "baseline":
			*m = 3
		case "bottom":
			*m = 4
		case "auto":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_TextAlignment) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "top"
	case 2:
		return "center"
	case 3:
		return "baseline"
	case 4:
		return "bottom"
	case 5:
		return "auto"
	}
	return ""
}

func (m ST_TextAlignment) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TextAlignment) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DisplacedByCustomXml byte

const (
	ST_DisplacedByCustomXmlUnset ST_DisplacedByCustomXml = 0
	ST_DisplacedByCustomXmlNext  ST_DisplacedByCustomXml = 1
	ST_DisplacedByCustomXmlPrev  ST_DisplacedByCustomXml = 2
)

func (e ST_DisplacedByCustomXml) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DisplacedByCustomXmlUnset:
		attr.Value = ""
	case ST_DisplacedByCustomXmlNext:
		attr.Value = "next"
	case ST_DisplacedByCustomXmlPrev:
		attr.Value = "prev"
	}
	return attr, nil
}

func (e *ST_DisplacedByCustomXml) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "next":
		*e = 1
	case "prev":
		*e = 2
	}
	return nil
}

func (m ST_DisplacedByCustomXml) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_DisplacedByCustomXml) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "next":
			*m = 1
		case "prev":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_DisplacedByCustomXml) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "next"
	case 2:
		return "prev"
	}
	return ""
}

func (m ST_DisplacedByCustomXml) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_DisplacedByCustomXml) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_AnnotationVMerge byte

const (
	ST_AnnotationVMergeUnset ST_AnnotationVMerge = 0
	ST_AnnotationVMergeCont  ST_AnnotationVMerge = 1
	ST_AnnotationVMergeRest  ST_AnnotationVMerge = 2
)

func (e ST_AnnotationVMerge) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_AnnotationVMergeUnset:
		attr.Value = ""
	case ST_AnnotationVMergeCont:
		attr.Value = "cont"
	case ST_AnnotationVMergeRest:
		attr.Value = "rest"
	}
	return attr, nil
}

func (e *ST_AnnotationVMerge) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "cont":
		*e = 1
	case "rest":
		*e = 2
	}
	return nil
}

func (m ST_AnnotationVMerge) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_AnnotationVMerge) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "cont":
			*m = 1
		case "rest":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_AnnotationVMerge) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "cont"
	case 2:
		return "rest"
	}
	return ""
}

func (m ST_AnnotationVMerge) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_AnnotationVMerge) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextboxTightWrap byte

const (
	ST_TextboxTightWrapUnset            ST_TextboxTightWrap = 0
	ST_TextboxTightWrapNone             ST_TextboxTightWrap = 1
	ST_TextboxTightWrapAllLines         ST_TextboxTightWrap = 2
	ST_TextboxTightWrapFirstAndLastLine ST_TextboxTightWrap = 3
	ST_TextboxTightWrapFirstLineOnly    ST_TextboxTightWrap = 4
	ST_TextboxTightWrapLastLineOnly     ST_TextboxTightWrap = 5
)

func (e ST_TextboxTightWrap) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextboxTightWrapUnset:
		attr.Value = ""
	case ST_TextboxTightWrapNone:
		attr.Value = "none"
	case ST_TextboxTightWrapAllLines:
		attr.Value = "allLines"
	case ST_TextboxTightWrapFirstAndLastLine:
		attr.Value = "firstAndLastLine"
	case ST_TextboxTightWrapFirstLineOnly:
		attr.Value = "firstLineOnly"
	case ST_TextboxTightWrapLastLineOnly:
		attr.Value = "lastLineOnly"
	}
	return attr, nil
}

func (e *ST_TextboxTightWrap) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "allLines":
		*e = 2
	case "firstAndLastLine":
		*e = 3
	case "firstLineOnly":
		*e = 4
	case "lastLineOnly":
		*e = 5
	}
	return nil
}

func (m ST_TextboxTightWrap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TextboxTightWrap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "allLines":
			*m = 2
		case "firstAndLastLine":
			*m = 3
		case "firstLineOnly":
			*m = 4
		case "lastLineOnly":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_TextboxTightWrap) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "allLines"
	case 3:
		return "firstAndLastLine"
	case 4:
		return "firstLineOnly"
	case 5:
		return "lastLineOnly"
	}
	return ""
}

func (m ST_TextboxTightWrap) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TextboxTightWrap) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ObjectDrawAspect byte

const (
	ST_ObjectDrawAspectUnset   ST_ObjectDrawAspect = 0
	ST_ObjectDrawAspectContent ST_ObjectDrawAspect = 1
	ST_ObjectDrawAspectIcon    ST_ObjectDrawAspect = 2
)

func (e ST_ObjectDrawAspect) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ObjectDrawAspectUnset:
		attr.Value = ""
	case ST_ObjectDrawAspectContent:
		attr.Value = "content"
	case ST_ObjectDrawAspectIcon:
		attr.Value = "icon"
	}
	return attr, nil
}

func (e *ST_ObjectDrawAspect) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "content":
		*e = 1
	case "icon":
		*e = 2
	}
	return nil
}

func (m ST_ObjectDrawAspect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ObjectDrawAspect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "content":
			*m = 1
		case "icon":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ObjectDrawAspect) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "content"
	case 2:
		return "icon"
	}
	return ""
}

func (m ST_ObjectDrawAspect) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ObjectDrawAspect) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ObjectUpdateMode byte

const (
	ST_ObjectUpdateModeUnset  ST_ObjectUpdateMode = 0
	ST_ObjectUpdateModeAlways ST_ObjectUpdateMode = 1
	ST_ObjectUpdateModeOnCall ST_ObjectUpdateMode = 2
)

func (e ST_ObjectUpdateMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ObjectUpdateModeUnset:
		attr.Value = ""
	case ST_ObjectUpdateModeAlways:
		attr.Value = "always"
	case ST_ObjectUpdateModeOnCall:
		attr.Value = "onCall"
	}
	return attr, nil
}

func (e *ST_ObjectUpdateMode) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "always":
		*e = 1
	case "onCall":
		*e = 2
	}
	return nil
}

func (m ST_ObjectUpdateMode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ObjectUpdateMode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "always":
			*m = 1
		case "onCall":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ObjectUpdateMode) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "always"
	case 2:
		return "onCall"
	}
	return ""
}

func (m ST_ObjectUpdateMode) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ObjectUpdateMode) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_FldCharType byte

const (
	ST_FldCharTypeUnset    ST_FldCharType = 0
	ST_FldCharTypeBegin    ST_FldCharType = 1
	ST_FldCharTypeSeparate ST_FldCharType = 2
	ST_FldCharTypeEnd      ST_FldCharType = 3
)

func (e ST_FldCharType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FldCharTypeUnset:
		attr.Value = ""
	case ST_FldCharTypeBegin:
		attr.Value = "begin"
	case ST_FldCharTypeSeparate:
		attr.Value = "separate"
	case ST_FldCharTypeEnd:
		attr.Value = "end"
	}
	return attr, nil
}

func (e *ST_FldCharType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "begin":
		*e = 1
	case "separate":
		*e = 2
	case "end":
		*e = 3
	}
	return nil
}

func (m ST_FldCharType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_FldCharType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "begin":
			*m = 1
		case "separate":
			*m = 2
		case "end":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_FldCharType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "begin"
	case 2:
		return "separate"
	case 3:
		return "end"
	}
	return ""
}

func (m ST_FldCharType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FldCharType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_InfoTextType byte

const (
	ST_InfoTextTypeUnset    ST_InfoTextType = 0
	ST_InfoTextTypeText     ST_InfoTextType = 1
	ST_InfoTextTypeAutoText ST_InfoTextType = 2
)

func (e ST_InfoTextType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_InfoTextTypeUnset:
		attr.Value = ""
	case ST_InfoTextTypeText:
		attr.Value = "text"
	case ST_InfoTextTypeAutoText:
		attr.Value = "autoText"
	}
	return attr, nil
}

func (e *ST_InfoTextType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "text":
		*e = 1
	case "autoText":
		*e = 2
	}
	return nil
}

func (m ST_InfoTextType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_InfoTextType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "text":
			*m = 1
		case "autoText":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_InfoTextType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "text"
	case 2:
		return "autoText"
	}
	return ""
}

func (m ST_InfoTextType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_InfoTextType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_FFTextType byte

const (
	ST_FFTextTypeUnset       ST_FFTextType = 0
	ST_FFTextTypeRegular     ST_FFTextType = 1
	ST_FFTextTypeNumber      ST_FFTextType = 2
	ST_FFTextTypeDate        ST_FFTextType = 3
	ST_FFTextTypeCurrentTime ST_FFTextType = 4
	ST_FFTextTypeCurrentDate ST_FFTextType = 5
	ST_FFTextTypeCalculated  ST_FFTextType = 6
)

func (e ST_FFTextType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FFTextTypeUnset:
		attr.Value = ""
	case ST_FFTextTypeRegular:
		attr.Value = "regular"
	case ST_FFTextTypeNumber:
		attr.Value = "number"
	case ST_FFTextTypeDate:
		attr.Value = "date"
	case ST_FFTextTypeCurrentTime:
		attr.Value = "currentTime"
	case ST_FFTextTypeCurrentDate:
		attr.Value = "currentDate"
	case ST_FFTextTypeCalculated:
		attr.Value = "calculated"
	}
	return attr, nil
}

func (e *ST_FFTextType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "regular":
		*e = 1
	case "number":
		*e = 2
	case "date":
		*e = 3
	case "currentTime":
		*e = 4
	case "currentDate":
		*e = 5
	case "calculated":
		*e = 6
	}
	return nil
}

func (m ST_FFTextType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_FFTextType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "regular":
			*m = 1
		case "number":
			*m = 2
		case "date":
			*m = 3
		case "currentTime":
			*m = 4
		case "currentDate":
			*m = 5
		case "calculated":
			*m = 6
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_FFTextType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "regular"
	case 2:
		return "number"
	case 3:
		return "date"
	case 4:
		return "currentTime"
	case 5:
		return "currentDate"
	case 6:
		return "calculated"
	}
	return ""
}

func (m ST_FFTextType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FFTextType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_SectionMark byte

const (
	ST_SectionMarkUnset      ST_SectionMark = 0
	ST_SectionMarkNextPage   ST_SectionMark = 1
	ST_SectionMarkNextColumn ST_SectionMark = 2
	ST_SectionMarkContinuous ST_SectionMark = 3
	ST_SectionMarkEvenPage   ST_SectionMark = 4
	ST_SectionMarkOddPage    ST_SectionMark = 5
)

func (e ST_SectionMark) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_SectionMarkUnset:
		attr.Value = ""
	case ST_SectionMarkNextPage:
		attr.Value = "nextPage"
	case ST_SectionMarkNextColumn:
		attr.Value = "nextColumn"
	case ST_SectionMarkContinuous:
		attr.Value = "continuous"
	case ST_SectionMarkEvenPage:
		attr.Value = "evenPage"
	case ST_SectionMarkOddPage:
		attr.Value = "oddPage"
	}
	return attr, nil
}

func (e *ST_SectionMark) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "nextPage":
		*e = 1
	case "nextColumn":
		*e = 2
	case "continuous":
		*e = 3
	case "evenPage":
		*e = 4
	case "oddPage":
		*e = 5
	}
	return nil
}

func (m ST_SectionMark) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_SectionMark) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "nextPage":
			*m = 1
		case "nextColumn":
			*m = 2
		case "continuous":
			*m = 3
		case "evenPage":
			*m = 4
		case "oddPage":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_SectionMark) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "nextPage"
	case 2:
		return "nextColumn"
	case 3:
		return "continuous"
	case 4:
		return "evenPage"
	case 5:
		return "oddPage"
	}
	return ""
}

func (m ST_SectionMark) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_SectionMark) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_NumberFormat byte

const (
	ST_NumberFormatUnset                        ST_NumberFormat = 0
	ST_NumberFormatDecimal                      ST_NumberFormat = 1
	ST_NumberFormatUpperRoman                   ST_NumberFormat = 2
	ST_NumberFormatLowerRoman                   ST_NumberFormat = 3
	ST_NumberFormatUpperLetter                  ST_NumberFormat = 4
	ST_NumberFormatLowerLetter                  ST_NumberFormat = 5
	ST_NumberFormatOrdinal                      ST_NumberFormat = 6
	ST_NumberFormatCardinalText                 ST_NumberFormat = 7
	ST_NumberFormatOrdinalText                  ST_NumberFormat = 8
	ST_NumberFormatHex                          ST_NumberFormat = 9
	ST_NumberFormatChicago                      ST_NumberFormat = 10
	ST_NumberFormatIdeographDigital             ST_NumberFormat = 11
	ST_NumberFormatJapaneseCounting             ST_NumberFormat = 12
	ST_NumberFormatAiueo                        ST_NumberFormat = 13
	ST_NumberFormatIroha                        ST_NumberFormat = 14
	ST_NumberFormatDecimalFullWidth             ST_NumberFormat = 15
	ST_NumberFormatDecimalHalfWidth             ST_NumberFormat = 16
	ST_NumberFormatJapaneseLegal                ST_NumberFormat = 17
	ST_NumberFormatJapaneseDigitalTenThousand   ST_NumberFormat = 18
	ST_NumberFormatDecimalEnclosedCircle        ST_NumberFormat = 19
	ST_NumberFormatDecimalFullWidth2            ST_NumberFormat = 20
	ST_NumberFormatAiueoFullWidth               ST_NumberFormat = 21
	ST_NumberFormatIrohaFullWidth               ST_NumberFormat = 22
	ST_NumberFormatDecimalZero                  ST_NumberFormat = 23
	ST_NumberFormatBullet                       ST_NumberFormat = 24
	ST_NumberFormatGanada                       ST_NumberFormat = 25
	ST_NumberFormatChosung                      ST_NumberFormat = 26
	ST_NumberFormatDecimalEnclosedFullstop      ST_NumberFormat = 27
	ST_NumberFormatDecimalEnclosedParen         ST_NumberFormat = 28
	ST_NumberFormatDecimalEnclosedCircleChinese ST_NumberFormat = 29
	ST_NumberFormatIdeographEnclosedCircle      ST_NumberFormat = 30
	ST_NumberFormatIdeographTraditional         ST_NumberFormat = 31
	ST_NumberFormatIdeographZodiac              ST_NumberFormat = 32
	ST_NumberFormatIdeographZodiacTraditional   ST_NumberFormat = 33
	ST_NumberFormatTaiwaneseCounting            ST_NumberFormat = 34
	ST_NumberFormatIdeographLegalTraditional    ST_NumberFormat = 35
	ST_NumberFormatTaiwaneseCountingThousand    ST_NumberFormat = 36
	ST_NumberFormatTaiwaneseDigital             ST_NumberFormat = 37
	ST_NumberFormatChineseCounting              ST_NumberFormat = 38
	ST_NumberFormatChineseLegalSimplified       ST_NumberFormat = 39
	ST_NumberFormatChineseCountingThousand      ST_NumberFormat = 40
	ST_NumberFormatKoreanDigital                ST_NumberFormat = 41
	ST_NumberFormatKoreanCounting               ST_NumberFormat = 42
	ST_NumberFormatKoreanLegal                  ST_NumberFormat = 43
	ST_NumberFormatKoreanDigital2               ST_NumberFormat = 44
	ST_NumberFormatVietnameseCounting           ST_NumberFormat = 45
	ST_NumberFormatRussianLower                 ST_NumberFormat = 46
	ST_NumberFormatRussianUpper                 ST_NumberFormat = 47
	ST_NumberFormatNone                         ST_NumberFormat = 48
	ST_NumberFormatNumberInDash                 ST_NumberFormat = 49
	ST_NumberFormatHebrew1                      ST_NumberFormat = 50
	ST_NumberFormatHebrew2                      ST_NumberFormat = 51
	ST_NumberFormatArabicAlpha                  ST_NumberFormat = 52
	ST_NumberFormatArabicAbjad                  ST_NumberFormat = 53
	ST_NumberFormatHindiVowels                  ST_NumberFormat = 54
	ST_NumberFormatHindiConsonants              ST_NumberFormat = 55
	ST_NumberFormatHindiNumbers                 ST_NumberFormat = 56
	ST_NumberFormatHindiCounting                ST_NumberFormat = 57
	ST_NumberFormatThaiLetters                  ST_NumberFormat = 58
	ST_NumberFormatThaiNumbers                  ST_NumberFormat = 59
	ST_NumberFormatThaiCounting                 ST_NumberFormat = 60
	ST_NumberFormatBahtText                     ST_NumberFormat = 61
	ST_NumberFormatDollarText                   ST_NumberFormat = 62
	ST_NumberFormatCustom                       ST_NumberFormat = 63
)

func (e ST_NumberFormat) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_NumberFormatUnset:
		attr.Value = ""
	case ST_NumberFormatDecimal:
		attr.Value = "decimal"
	case ST_NumberFormatUpperRoman:
		attr.Value = "upperRoman"
	case ST_NumberFormatLowerRoman:
		attr.Value = "lowerRoman"
	case ST_NumberFormatUpperLetter:
		attr.Value = "upperLetter"
	case ST_NumberFormatLowerLetter:
		attr.Value = "lowerLetter"
	case ST_NumberFormatOrdinal:
		attr.Value = "ordinal"
	case ST_NumberFormatCardinalText:
		attr.Value = "cardinalText"
	case ST_NumberFormatOrdinalText:
		attr.Value = "ordinalText"
	case ST_NumberFormatHex:
		attr.Value = "hex"
	case ST_NumberFormatChicago:
		attr.Value = "chicago"
	case ST_NumberFormatIdeographDigital:
		attr.Value = "ideographDigital"
	case ST_NumberFormatJapaneseCounting:
		attr.Value = "japaneseCounting"
	case ST_NumberFormatAiueo:
		attr.Value = "aiueo"
	case ST_NumberFormatIroha:
		attr.Value = "iroha"
	case ST_NumberFormatDecimalFullWidth:
		attr.Value = "decimalFullWidth"
	case ST_NumberFormatDecimalHalfWidth:
		attr.Value = "decimalHalfWidth"
	case ST_NumberFormatJapaneseLegal:
		attr.Value = "japaneseLegal"
	case ST_NumberFormatJapaneseDigitalTenThousand:
		attr.Value = "japaneseDigitalTenThousand"
	case ST_NumberFormatDecimalEnclosedCircle:
		attr.Value = "decimalEnclosedCircle"
	case ST_NumberFormatDecimalFullWidth2:
		attr.Value = "decimalFullWidth2"
	case ST_NumberFormatAiueoFullWidth:
		attr.Value = "aiueoFullWidth"
	case ST_NumberFormatIrohaFullWidth:
		attr.Value = "irohaFullWidth"
	case ST_NumberFormatDecimalZero:
		attr.Value = "decimalZero"
	case ST_NumberFormatBullet:
		attr.Value = "bullet"
	case ST_NumberFormatGanada:
		attr.Value = "ganada"
	case ST_NumberFormatChosung:
		attr.Value = "chosung"
	case ST_NumberFormatDecimalEnclosedFullstop:
		attr.Value = "decimalEnclosedFullstop"
	case ST_NumberFormatDecimalEnclosedParen:
		attr.Value = "decimalEnclosedParen"
	case ST_NumberFormatDecimalEnclosedCircleChinese:
		attr.Value = "decimalEnclosedCircleChinese"
	case ST_NumberFormatIdeographEnclosedCircle:
		attr.Value = "ideographEnclosedCircle"
	case ST_NumberFormatIdeographTraditional:
		attr.Value = "ideographTraditional"
	case ST_NumberFormatIdeographZodiac:
		attr.Value = "ideographZodiac"
	case ST_NumberFormatIdeographZodiacTraditional:
		attr.Value = "ideographZodiacTraditional"
	case ST_NumberFormatTaiwaneseCounting:
		attr.Value = "taiwaneseCounting"
	case ST_NumberFormatIdeographLegalTraditional:
		attr.Value = "ideographLegalTraditional"
	case ST_NumberFormatTaiwaneseCountingThousand:
		attr.Value = "taiwaneseCountingThousand"
	case ST_NumberFormatTaiwaneseDigital:
		attr.Value = "taiwaneseDigital"
	case ST_NumberFormatChineseCounting:
		attr.Value = "chineseCounting"
	case ST_NumberFormatChineseLegalSimplified:
		attr.Value = "chineseLegalSimplified"
	case ST_NumberFormatChineseCountingThousand:
		attr.Value = "chineseCountingThousand"
	case ST_NumberFormatKoreanDigital:
		attr.Value = "koreanDigital"
	case ST_NumberFormatKoreanCounting:
		attr.Value = "koreanCounting"
	case ST_NumberFormatKoreanLegal:
		attr.Value = "koreanLegal"
	case ST_NumberFormatKoreanDigital2:
		attr.Value = "koreanDigital2"
	case ST_NumberFormatVietnameseCounting:
		attr.Value = "vietnameseCounting"
	case ST_NumberFormatRussianLower:
		attr.Value = "russianLower"
	case ST_NumberFormatRussianUpper:
		attr.Value = "russianUpper"
	case ST_NumberFormatNone:
		attr.Value = "none"
	case ST_NumberFormatNumberInDash:
		attr.Value = "numberInDash"
	case ST_NumberFormatHebrew1:
		attr.Value = "hebrew1"
	case ST_NumberFormatHebrew2:
		attr.Value = "hebrew2"
	case ST_NumberFormatArabicAlpha:
		attr.Value = "arabicAlpha"
	case ST_NumberFormatArabicAbjad:
		attr.Value = "arabicAbjad"
	case ST_NumberFormatHindiVowels:
		attr.Value = "hindiVowels"
	case ST_NumberFormatHindiConsonants:
		attr.Value = "hindiConsonants"
	case ST_NumberFormatHindiNumbers:
		attr.Value = "hindiNumbers"
	case ST_NumberFormatHindiCounting:
		attr.Value = "hindiCounting"
	case ST_NumberFormatThaiLetters:
		attr.Value = "thaiLetters"
	case ST_NumberFormatThaiNumbers:
		attr.Value = "thaiNumbers"
	case ST_NumberFormatThaiCounting:
		attr.Value = "thaiCounting"
	case ST_NumberFormatBahtText:
		attr.Value = "bahtText"
	case ST_NumberFormatDollarText:
		attr.Value = "dollarText"
	case ST_NumberFormatCustom:
		attr.Value = "custom"
	}
	return attr, nil
}

func (e *ST_NumberFormat) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "decimal":
		*e = 1
	case "upperRoman":
		*e = 2
	case "lowerRoman":
		*e = 3
	case "upperLetter":
		*e = 4
	case "lowerLetter":
		*e = 5
	case "ordinal":
		*e = 6
	case "cardinalText":
		*e = 7
	case "ordinalText":
		*e = 8
	case "hex":
		*e = 9
	case "chicago":
		*e = 10
	case "ideographDigital":
		*e = 11
	case "japaneseCounting":
		*e = 12
	case "aiueo":
		*e = 13
	case "iroha":
		*e = 14
	case "decimalFullWidth":
		*e = 15
	case "decimalHalfWidth":
		*e = 16
	case "japaneseLegal":
		*e = 17
	case "japaneseDigitalTenThousand":
		*e = 18
	case "decimalEnclosedCircle":
		*e = 19
	case "decimalFullWidth2":
		*e = 20
	case "aiueoFullWidth":
		*e = 21
	case "irohaFullWidth":
		*e = 22
	case "decimalZero":
		*e = 23
	case "bullet":
		*e = 24
	case "ganada":
		*e = 25
	case "chosung":
		*e = 26
	case "decimalEnclosedFullstop":
		*e = 27
	case "decimalEnclosedParen":
		*e = 28
	case "decimalEnclosedCircleChinese":
		*e = 29
	case "ideographEnclosedCircle":
		*e = 30
	case "ideographTraditional":
		*e = 31
	case "ideographZodiac":
		*e = 32
	case "ideographZodiacTraditional":
		*e = 33
	case "taiwaneseCounting":
		*e = 34
	case "ideographLegalTraditional":
		*e = 35
	case "taiwaneseCountingThousand":
		*e = 36
	case "taiwaneseDigital":
		*e = 37
	case "chineseCounting":
		*e = 38
	case "chineseLegalSimplified":
		*e = 39
	case "chineseCountingThousand":
		*e = 40
	case "koreanDigital":
		*e = 41
	case "koreanCounting":
		*e = 42
	case "koreanLegal":
		*e = 43
	case "koreanDigital2":
		*e = 44
	case "vietnameseCounting":
		*e = 45
	case "russianLower":
		*e = 46
	case "russianUpper":
		*e = 47
	case "none":
		*e = 48
	case "numberInDash":
		*e = 49
	case "hebrew1":
		*e = 50
	case "hebrew2":
		*e = 51
	case "arabicAlpha":
		*e = 52
	case "arabicAbjad":
		*e = 53
	case "hindiVowels":
		*e = 54
	case "hindiConsonants":
		*e = 55
	case "hindiNumbers":
		*e = 56
	case "hindiCounting":
		*e = 57
	case "thaiLetters":
		*e = 58
	case "thaiNumbers":
		*e = 59
	case "thaiCounting":
		*e = 60
	case "bahtText":
		*e = 61
	case "dollarText":
		*e = 62
	case "custom":
		*e = 63
	}
	return nil
}

func (m ST_NumberFormat) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_NumberFormat) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "decimal":
			*m = 1
		case "upperRoman":
			*m = 2
		case "lowerRoman":
			*m = 3
		case "upperLetter":
			*m = 4
		case "lowerLetter":
			*m = 5
		case "ordinal":
			*m = 6
		case "cardinalText":
			*m = 7
		case "ordinalText":
			*m = 8
		case "hex":
			*m = 9
		case "chicago":
			*m = 10
		case "ideographDigital":
			*m = 11
		case "japaneseCounting":
			*m = 12
		case "aiueo":
			*m = 13
		case "iroha":
			*m = 14
		case "decimalFullWidth":
			*m = 15
		case "decimalHalfWidth":
			*m = 16
		case "japaneseLegal":
			*m = 17
		case "japaneseDigitalTenThousand":
			*m = 18
		case "decimalEnclosedCircle":
			*m = 19
		case "decimalFullWidth2":
			*m = 20
		case "aiueoFullWidth":
			*m = 21
		case "irohaFullWidth":
			*m = 22
		case "decimalZero":
			*m = 23
		case "bullet":
			*m = 24
		case "ganada":
			*m = 25
		case "chosung":
			*m = 26
		case "decimalEnclosedFullstop":
			*m = 27
		case "decimalEnclosedParen":
			*m = 28
		case "decimalEnclosedCircleChinese":
			*m = 29
		case "ideographEnclosedCircle":
			*m = 30
		case "ideographTraditional":
			*m = 31
		case "ideographZodiac":
			*m = 32
		case "ideographZodiacTraditional":
			*m = 33
		case "taiwaneseCounting":
			*m = 34
		case "ideographLegalTraditional":
			*m = 35
		case "taiwaneseCountingThousand":
			*m = 36
		case "taiwaneseDigital":
			*m = 37
		case "chineseCounting":
			*m = 38
		case "chineseLegalSimplified":
			*m = 39
		case "chineseCountingThousand":
			*m = 40
		case "koreanDigital":
			*m = 41
		case "koreanCounting":
			*m = 42
		case "koreanLegal":
			*m = 43
		case "koreanDigital2":
			*m = 44
		case "vietnameseCounting":
			*m = 45
		case "russianLower":
			*m = 46
		case "russianUpper":
			*m = 47
		case "none":
			*m = 48
		case "numberInDash":
			*m = 49
		case "hebrew1":
			*m = 50
		case "hebrew2":
			*m = 51
		case "arabicAlpha":
			*m = 52
		case "arabicAbjad":
			*m = 53
		case "hindiVowels":
			*m = 54
		case "hindiConsonants":
			*m = 55
		case "hindiNumbers":
			*m = 56
		case "hindiCounting":
			*m = 57
		case "thaiLetters":
			*m = 58
		case "thaiNumbers":
			*m = 59
		case "thaiCounting":
			*m = 60
		case "bahtText":
			*m = 61
		case "dollarText":
			*m = 62
		case "custom":
			*m = 63
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_NumberFormat) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "decimal"
	case 2:
		return "upperRoman"
	case 3:
		return "lowerRoman"
	case 4:
		return "upperLetter"
	case 5:
		return "lowerLetter"
	case 6:
		return "ordinal"
	case 7:
		return "cardinalText"
	case 8:
		return "ordinalText"
	case 9:
		return "hex"
	case 10:
		return "chicago"
	case 11:
		return "ideographDigital"
	case 12:
		return "japaneseCounting"
	case 13:
		return "aiueo"
	case 14:
		return "iroha"
	case 15:
		return "decimalFullWidth"
	case 16:
		return "decimalHalfWidth"
	case 17:
		return "japaneseLegal"
	case 18:
		return "japaneseDigitalTenThousand"
	case 19:
		return "decimalEnclosedCircle"
	case 20:
		return "decimalFullWidth2"
	case 21:
		return "aiueoFullWidth"
	case 22:
		return "irohaFullWidth"
	case 23:
		return "decimalZero"
	case 24:
		return "bullet"
	case 25:
		return "ganada"
	case 26:
		return "chosung"
	case 27:
		return "decimalEnclosedFullstop"
	case 28:
		return "decimalEnclosedParen"
	case 29:
		return "decimalEnclosedCircleChinese"
	case 30:
		return "ideographEnclosedCircle"
	case 31:
		return "ideographTraditional"
	case 32:
		return "ideographZodiac"
	case 33:
		return "ideographZodiacTraditional"
	case 34:
		return "taiwaneseCounting"
	case 35:
		return "ideographLegalTraditional"
	case 36:
		return "taiwaneseCountingThousand"
	case 37:
		return "taiwaneseDigital"
	case 38:
		return "chineseCounting"
	case 39:
		return "chineseLegalSimplified"
	case 40:
		return "chineseCountingThousand"
	case 41:
		return "koreanDigital"
	case 42:
		return "koreanCounting"
	case 43:
		return "koreanLegal"
	case 44:
		return "koreanDigital2"
	case 45:
		return "vietnameseCounting"
	case 46:
		return "russianLower"
	case 47:
		return "russianUpper"
	case 48:
		return "none"
	case 49:
		return "numberInDash"
	case 50:
		return "hebrew1"
	case 51:
		return "hebrew2"
	case 52:
		return "arabicAlpha"
	case 53:
		return "arabicAbjad"
	case 54:
		return "hindiVowels"
	case 55:
		return "hindiConsonants"
	case 56:
		return "hindiNumbers"
	case 57:
		return "hindiCounting"
	case 58:
		return "thaiLetters"
	case 59:
		return "thaiNumbers"
	case 60:
		return "thaiCounting"
	case 61:
		return "bahtText"
	case 62:
		return "dollarText"
	case 63:
		return "custom"
	}
	return ""
}

func (m ST_NumberFormat) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_NumberFormat) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PageOrientation byte

const (
	ST_PageOrientationUnset     ST_PageOrientation = 0
	ST_PageOrientationPortrait  ST_PageOrientation = 1
	ST_PageOrientationLandscape ST_PageOrientation = 2
)

func (e ST_PageOrientation) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PageOrientationUnset:
		attr.Value = ""
	case ST_PageOrientationPortrait:
		attr.Value = "portrait"
	case ST_PageOrientationLandscape:
		attr.Value = "landscape"
	}
	return attr, nil
}

func (e *ST_PageOrientation) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "portrait":
		*e = 1
	case "landscape":
		*e = 2
	}
	return nil
}

func (m ST_PageOrientation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_PageOrientation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "portrait":
			*m = 1
		case "landscape":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_PageOrientation) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "portrait"
	case 2:
		return "landscape"
	}
	return ""
}

func (m ST_PageOrientation) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PageOrientation) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PageBorderZOrder byte

const (
	ST_PageBorderZOrderUnset ST_PageBorderZOrder = 0
	ST_PageBorderZOrderFront ST_PageBorderZOrder = 1
	ST_PageBorderZOrderBack  ST_PageBorderZOrder = 2
)

func (e ST_PageBorderZOrder) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PageBorderZOrderUnset:
		attr.Value = ""
	case ST_PageBorderZOrderFront:
		attr.Value = "front"
	case ST_PageBorderZOrderBack:
		attr.Value = "back"
	}
	return attr, nil
}

func (e *ST_PageBorderZOrder) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "front":
		*e = 1
	case "back":
		*e = 2
	}
	return nil
}

func (m ST_PageBorderZOrder) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_PageBorderZOrder) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "front":
			*m = 1
		case "back":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_PageBorderZOrder) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "front"
	case 2:
		return "back"
	}
	return ""
}

func (m ST_PageBorderZOrder) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PageBorderZOrder) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PageBorderDisplay byte

const (
	ST_PageBorderDisplayUnset        ST_PageBorderDisplay = 0
	ST_PageBorderDisplayAllPages     ST_PageBorderDisplay = 1
	ST_PageBorderDisplayFirstPage    ST_PageBorderDisplay = 2
	ST_PageBorderDisplayNotFirstPage ST_PageBorderDisplay = 3
)

func (e ST_PageBorderDisplay) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PageBorderDisplayUnset:
		attr.Value = ""
	case ST_PageBorderDisplayAllPages:
		attr.Value = "allPages"
	case ST_PageBorderDisplayFirstPage:
		attr.Value = "firstPage"
	case ST_PageBorderDisplayNotFirstPage:
		attr.Value = "notFirstPage"
	}
	return attr, nil
}

func (e *ST_PageBorderDisplay) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "allPages":
		*e = 1
	case "firstPage":
		*e = 2
	case "notFirstPage":
		*e = 3
	}
	return nil
}

func (m ST_PageBorderDisplay) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_PageBorderDisplay) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "allPages":
			*m = 1
		case "firstPage":
			*m = 2
		case "notFirstPage":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_PageBorderDisplay) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "allPages"
	case 2:
		return "firstPage"
	case 3:
		return "notFirstPage"
	}
	return ""
}

func (m ST_PageBorderDisplay) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PageBorderDisplay) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PageBorderOffset byte

const (
	ST_PageBorderOffsetUnset ST_PageBorderOffset = 0
	ST_PageBorderOffsetPage  ST_PageBorderOffset = 1
	ST_PageBorderOffsetText  ST_PageBorderOffset = 2
)

func (e ST_PageBorderOffset) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PageBorderOffsetUnset:
		attr.Value = ""
	case ST_PageBorderOffsetPage:
		attr.Value = "page"
	case ST_PageBorderOffsetText:
		attr.Value = "text"
	}
	return attr, nil
}

func (e *ST_PageBorderOffset) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "page":
		*e = 1
	case "text":
		*e = 2
	}
	return nil
}

func (m ST_PageBorderOffset) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_PageBorderOffset) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "page":
			*m = 1
		case "text":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_PageBorderOffset) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "page"
	case 2:
		return "text"
	}
	return ""
}

func (m ST_PageBorderOffset) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PageBorderOffset) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ChapterSep byte

const (
	ST_ChapterSepUnset  ST_ChapterSep = 0
	ST_ChapterSepHyphen ST_ChapterSep = 1
	ST_ChapterSepPeriod ST_ChapterSep = 2
	ST_ChapterSepColon  ST_ChapterSep = 3
	ST_ChapterSepEmDash ST_ChapterSep = 4
	ST_ChapterSepEnDash ST_ChapterSep = 5
)

func (e ST_ChapterSep) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ChapterSepUnset:
		attr.Value = ""
	case ST_ChapterSepHyphen:
		attr.Value = "hyphen"
	case ST_ChapterSepPeriod:
		attr.Value = "period"
	case ST_ChapterSepColon:
		attr.Value = "colon"
	case ST_ChapterSepEmDash:
		attr.Value = "emDash"
	case ST_ChapterSepEnDash:
		attr.Value = "enDash"
	}
	return attr, nil
}

func (e *ST_ChapterSep) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "hyphen":
		*e = 1
	case "period":
		*e = 2
	case "colon":
		*e = 3
	case "emDash":
		*e = 4
	case "enDash":
		*e = 5
	}
	return nil
}

func (m ST_ChapterSep) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ChapterSep) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "hyphen":
			*m = 1
		case "period":
			*m = 2
		case "colon":
			*m = 3
		case "emDash":
			*m = 4
		case "enDash":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ChapterSep) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "hyphen"
	case 2:
		return "period"
	case 3:
		return "colon"
	case 4:
		return "emDash"
	case 5:
		return "enDash"
	}
	return ""
}

func (m ST_ChapterSep) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ChapterSep) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_LineNumberRestart byte

const (
	ST_LineNumberRestartUnset      ST_LineNumberRestart = 0
	ST_LineNumberRestartNewPage    ST_LineNumberRestart = 1
	ST_LineNumberRestartNewSection ST_LineNumberRestart = 2
	ST_LineNumberRestartContinuous ST_LineNumberRestart = 3
)

func (e ST_LineNumberRestart) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_LineNumberRestartUnset:
		attr.Value = ""
	case ST_LineNumberRestartNewPage:
		attr.Value = "newPage"
	case ST_LineNumberRestartNewSection:
		attr.Value = "newSection"
	case ST_LineNumberRestartContinuous:
		attr.Value = "continuous"
	}
	return attr, nil
}

func (e *ST_LineNumberRestart) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "newPage":
		*e = 1
	case "newSection":
		*e = 2
	case "continuous":
		*e = 3
	}
	return nil
}

func (m ST_LineNumberRestart) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_LineNumberRestart) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "newPage":
			*m = 1
		case "newSection":
			*m = 2
		case "continuous":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_LineNumberRestart) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "newPage"
	case 2:
		return "newSection"
	case 3:
		return "continuous"
	}
	return ""
}

func (m ST_LineNumberRestart) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_LineNumberRestart) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_VerticalJc byte

const (
	ST_VerticalJcUnset  ST_VerticalJc = 0
	ST_VerticalJcTop    ST_VerticalJc = 1
	ST_VerticalJcCenter ST_VerticalJc = 2
	ST_VerticalJcBoth   ST_VerticalJc = 3
	ST_VerticalJcBottom ST_VerticalJc = 4
)

func (e ST_VerticalJc) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_VerticalJcUnset:
		attr.Value = ""
	case ST_VerticalJcTop:
		attr.Value = "top"
	case ST_VerticalJcCenter:
		attr.Value = "center"
	case ST_VerticalJcBoth:
		attr.Value = "both"
	case ST_VerticalJcBottom:
		attr.Value = "bottom"
	}
	return attr, nil
}

func (e *ST_VerticalJc) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "top":
		*e = 1
	case "center":
		*e = 2
	case "both":
		*e = 3
	case "bottom":
		*e = 4
	}
	return nil
}

func (m ST_VerticalJc) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_VerticalJc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "top":
			*m = 1
		case "center":
			*m = 2
		case "both":
			*m = 3
		case "bottom":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_VerticalJc) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "top"
	case 2:
		return "center"
	case 3:
		return "both"
	case 4:
		return "bottom"
	}
	return ""
}

func (m ST_VerticalJc) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_VerticalJc) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DocGrid byte

const (
	ST_DocGridUnset         ST_DocGrid = 0
	ST_DocGridDefault       ST_DocGrid = 1
	ST_DocGridLines         ST_DocGrid = 2
	ST_DocGridLinesAndChars ST_DocGrid = 3
	ST_DocGridSnapToChars   ST_DocGrid = 4
)

func (e ST_DocGrid) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DocGridUnset:
		attr.Value = ""
	case ST_DocGridDefault:
		attr.Value = "default"
	case ST_DocGridLines:
		attr.Value = "lines"
	case ST_DocGridLinesAndChars:
		attr.Value = "linesAndChars"
	case ST_DocGridSnapToChars:
		attr.Value = "snapToChars"
	}
	return attr, nil
}

func (e *ST_DocGrid) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "default":
		*e = 1
	case "lines":
		*e = 2
	case "linesAndChars":
		*e = 3
	case "snapToChars":
		*e = 4
	}
	return nil
}

func (m ST_DocGrid) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_DocGrid) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "default":
			*m = 1
		case "lines":
			*m = 2
		case "linesAndChars":
			*m = 3
		case "snapToChars":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_DocGrid) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "default"
	case 2:
		return "lines"
	case 3:
		return "linesAndChars"
	case 4:
		return "snapToChars"
	}
	return ""
}

func (m ST_DocGrid) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_DocGrid) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_HdrFtr byte

const (
	ST_HdrFtrUnset   ST_HdrFtr = 0
	ST_HdrFtrEven    ST_HdrFtr = 1
	ST_HdrFtrDefault ST_HdrFtr = 2
	ST_HdrFtrFirst   ST_HdrFtr = 3
)

func (e ST_HdrFtr) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_HdrFtrUnset:
		attr.Value = ""
	case ST_HdrFtrEven:
		attr.Value = "even"
	case ST_HdrFtrDefault:
		attr.Value = "default"
	case ST_HdrFtrFirst:
		attr.Value = "first"
	}
	return attr, nil
}

func (e *ST_HdrFtr) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "even":
		*e = 1
	case "default":
		*e = 2
	case "first":
		*e = 3
	}
	return nil
}

func (m ST_HdrFtr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_HdrFtr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "even":
			*m = 1
		case "default":
			*m = 2
		case "first":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_HdrFtr) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "even"
	case 2:
		return "default"
	case 3:
		return "first"
	}
	return ""
}

func (m ST_HdrFtr) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_HdrFtr) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_FtnEdn byte

const (
	ST_FtnEdnUnset                 ST_FtnEdn = 0
	ST_FtnEdnNormal                ST_FtnEdn = 1
	ST_FtnEdnSeparator             ST_FtnEdn = 2
	ST_FtnEdnContinuationSeparator ST_FtnEdn = 3
	ST_FtnEdnContinuationNotice    ST_FtnEdn = 4
)

func (e ST_FtnEdn) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FtnEdnUnset:
		attr.Value = ""
	case ST_FtnEdnNormal:
		attr.Value = "normal"
	case ST_FtnEdnSeparator:
		attr.Value = "separator"
	case ST_FtnEdnContinuationSeparator:
		attr.Value = "continuationSeparator"
	case ST_FtnEdnContinuationNotice:
		attr.Value = "continuationNotice"
	}
	return attr, nil
}

func (e *ST_FtnEdn) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "normal":
		*e = 1
	case "separator":
		*e = 2
	case "continuationSeparator":
		*e = 3
	case "continuationNotice":
		*e = 4
	}
	return nil
}

func (m ST_FtnEdn) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_FtnEdn) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "normal":
			*m = 1
		case "separator":
			*m = 2
		case "continuationSeparator":
			*m = 3
		case "continuationNotice":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_FtnEdn) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "normal"
	case 2:
		return "separator"
	case 3:
		return "continuationSeparator"
	case 4:
		return "continuationNotice"
	}
	return ""
}

func (m ST_FtnEdn) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FtnEdn) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_BrType byte

const (
	ST_BrTypeUnset        ST_BrType = 0
	ST_BrTypePage         ST_BrType = 1
	ST_BrTypeColumn       ST_BrType = 2
	ST_BrTypeTextWrapping ST_BrType = 3
)

func (e ST_BrType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_BrTypeUnset:
		attr.Value = ""
	case ST_BrTypePage:
		attr.Value = "page"
	case ST_BrTypeColumn:
		attr.Value = "column"
	case ST_BrTypeTextWrapping:
		attr.Value = "textWrapping"
	}
	return attr, nil
}

func (e *ST_BrType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "page":
		*e = 1
	case "column":
		*e = 2
	case "textWrapping":
		*e = 3
	}
	return nil
}

func (m ST_BrType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_BrType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "page":
			*m = 1
		case "column":
			*m = 2
		case "textWrapping":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_BrType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "page"
	case 2:
		return "column"
	case 3:
		return "textWrapping"
	}
	return ""
}

func (m ST_BrType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_BrType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_BrClear byte

const (
	ST_BrClearUnset ST_BrClear = 0
	ST_BrClearNone  ST_BrClear = 1
	ST_BrClearLeft  ST_BrClear = 2
	ST_BrClearRight ST_BrClear = 3
	ST_BrClearAll   ST_BrClear = 4
)

func (e ST_BrClear) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_BrClearUnset:
		attr.Value = ""
	case ST_BrClearNone:
		attr.Value = "none"
	case ST_BrClearLeft:
		attr.Value = "left"
	case ST_BrClearRight:
		attr.Value = "right"
	case ST_BrClearAll:
		attr.Value = "all"
	}
	return attr, nil
}

func (e *ST_BrClear) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "left":
		*e = 2
	case "right":
		*e = 3
	case "all":
		*e = 4
	}
	return nil
}

func (m ST_BrClear) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_BrClear) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "left":
			*m = 2
		case "right":
			*m = 3
		case "all":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_BrClear) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "left"
	case 3:
		return "right"
	case 4:
		return "all"
	}
	return ""
}

func (m ST_BrClear) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_BrClear) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PTabAlignment byte

const (
	ST_PTabAlignmentUnset  ST_PTabAlignment = 0
	ST_PTabAlignmentLeft   ST_PTabAlignment = 1
	ST_PTabAlignmentCenter ST_PTabAlignment = 2
	ST_PTabAlignmentRight  ST_PTabAlignment = 3
)

func (e ST_PTabAlignment) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PTabAlignmentUnset:
		attr.Value = ""
	case ST_PTabAlignmentLeft:
		attr.Value = "left"
	case ST_PTabAlignmentCenter:
		attr.Value = "center"
	case ST_PTabAlignmentRight:
		attr.Value = "right"
	}
	return attr, nil
}

func (e *ST_PTabAlignment) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "left":
		*e = 1
	case "center":
		*e = 2
	case "right":
		*e = 3
	}
	return nil
}

func (m ST_PTabAlignment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_PTabAlignment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "left":
			*m = 1
		case "center":
			*m = 2
		case "right":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_PTabAlignment) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "left"
	case 2:
		return "center"
	case 3:
		return "right"
	}
	return ""
}

func (m ST_PTabAlignment) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PTabAlignment) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PTabRelativeTo byte

const (
	ST_PTabRelativeToUnset  ST_PTabRelativeTo = 0
	ST_PTabRelativeToMargin ST_PTabRelativeTo = 1
	ST_PTabRelativeToIndent ST_PTabRelativeTo = 2
)

func (e ST_PTabRelativeTo) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PTabRelativeToUnset:
		attr.Value = ""
	case ST_PTabRelativeToMargin:
		attr.Value = "margin"
	case ST_PTabRelativeToIndent:
		attr.Value = "indent"
	}
	return attr, nil
}

func (e *ST_PTabRelativeTo) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "margin":
		*e = 1
	case "indent":
		*e = 2
	}
	return nil
}

func (m ST_PTabRelativeTo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_PTabRelativeTo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "margin":
			*m = 1
		case "indent":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_PTabRelativeTo) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "margin"
	case 2:
		return "indent"
	}
	return ""
}

func (m ST_PTabRelativeTo) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PTabRelativeTo) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PTabLeader byte

const (
	ST_PTabLeaderUnset      ST_PTabLeader = 0
	ST_PTabLeaderNone       ST_PTabLeader = 1
	ST_PTabLeaderDot        ST_PTabLeader = 2
	ST_PTabLeaderHyphen     ST_PTabLeader = 3
	ST_PTabLeaderUnderscore ST_PTabLeader = 4
	ST_PTabLeaderMiddleDot  ST_PTabLeader = 5
)

func (e ST_PTabLeader) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PTabLeaderUnset:
		attr.Value = ""
	case ST_PTabLeaderNone:
		attr.Value = "none"
	case ST_PTabLeaderDot:
		attr.Value = "dot"
	case ST_PTabLeaderHyphen:
		attr.Value = "hyphen"
	case ST_PTabLeaderUnderscore:
		attr.Value = "underscore"
	case ST_PTabLeaderMiddleDot:
		attr.Value = "middleDot"
	}
	return attr, nil
}

func (e *ST_PTabLeader) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "dot":
		*e = 2
	case "hyphen":
		*e = 3
	case "underscore":
		*e = 4
	case "middleDot":
		*e = 5
	}
	return nil
}

func (m ST_PTabLeader) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_PTabLeader) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "dot":
			*m = 2
		case "hyphen":
			*m = 3
		case "underscore":
			*m = 4
		case "middleDot":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_PTabLeader) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "dot"
	case 3:
		return "hyphen"
	case 4:
		return "underscore"
	case 5:
		return "middleDot"
	}
	return ""
}

func (m ST_PTabLeader) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PTabLeader) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ProofErr byte

const (
	ST_ProofErrUnset      ST_ProofErr = 0
	ST_ProofErrSpellStart ST_ProofErr = 1
	ST_ProofErrSpellEnd   ST_ProofErr = 2
	ST_ProofErrGramStart  ST_ProofErr = 3
	ST_ProofErrGramEnd    ST_ProofErr = 4
)

func (e ST_ProofErr) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ProofErrUnset:
		attr.Value = ""
	case ST_ProofErrSpellStart:
		attr.Value = "spellStart"
	case ST_ProofErrSpellEnd:
		attr.Value = "spellEnd"
	case ST_ProofErrGramStart:
		attr.Value = "gramStart"
	case ST_ProofErrGramEnd:
		attr.Value = "gramEnd"
	}
	return attr, nil
}

func (e *ST_ProofErr) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "spellStart":
		*e = 1
	case "spellEnd":
		*e = 2
	case "gramStart":
		*e = 3
	case "gramEnd":
		*e = 4
	}
	return nil
}

func (m ST_ProofErr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ProofErr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "spellStart":
			*m = 1
		case "spellEnd":
			*m = 2
		case "gramStart":
			*m = 3
		case "gramEnd":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ProofErr) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "spellStart"
	case 2:
		return "spellEnd"
	case 3:
		return "gramStart"
	case 4:
		return "gramEnd"
	}
	return ""
}

func (m ST_ProofErr) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ProofErr) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_EdGrp byte

const (
	ST_EdGrpUnset          ST_EdGrp = 0
	ST_EdGrpNone           ST_EdGrp = 1
	ST_EdGrpEveryone       ST_EdGrp = 2
	ST_EdGrpAdministrators ST_EdGrp = 3
	ST_EdGrpContributors   ST_EdGrp = 4
	ST_EdGrpEditors        ST_EdGrp = 5
	ST_EdGrpOwners         ST_EdGrp = 6
	ST_EdGrpCurrent        ST_EdGrp = 7
)

func (e ST_EdGrp) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_EdGrpUnset:
		attr.Value = ""
	case ST_EdGrpNone:
		attr.Value = "none"
	case ST_EdGrpEveryone:
		attr.Value = "everyone"
	case ST_EdGrpAdministrators:
		attr.Value = "administrators"
	case ST_EdGrpContributors:
		attr.Value = "contributors"
	case ST_EdGrpEditors:
		attr.Value = "editors"
	case ST_EdGrpOwners:
		attr.Value = "owners"
	case ST_EdGrpCurrent:
		attr.Value = "current"
	}
	return attr, nil
}

func (e *ST_EdGrp) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "everyone":
		*e = 2
	case "administrators":
		*e = 3
	case "contributors":
		*e = 4
	case "editors":
		*e = 5
	case "owners":
		*e = 6
	case "current":
		*e = 7
	}
	return nil
}

func (m ST_EdGrp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_EdGrp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "everyone":
			*m = 2
		case "administrators":
			*m = 3
		case "contributors":
			*m = 4
		case "editors":
			*m = 5
		case "owners":
			*m = 6
		case "current":
			*m = 7
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_EdGrp) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "everyone"
	case 3:
		return "administrators"
	case 4:
		return "contributors"
	case 5:
		return "editors"
	case 6:
		return "owners"
	case 7:
		return "current"
	}
	return ""
}

func (m ST_EdGrp) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_EdGrp) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Hint byte

const (
	ST_HintUnset    ST_Hint = 0
	ST_HintDefault  ST_Hint = 1
	ST_HintEastAsia ST_Hint = 2
)

func (e ST_Hint) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_HintUnset:
		attr.Value = ""
	case ST_HintDefault:
		attr.Value = "default"
	case ST_HintEastAsia:
		attr.Value = "eastAsia"
	}
	return attr, nil
}

func (e *ST_Hint) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "default":
		*e = 1
	case "eastAsia":
		*e = 2
	}
	return nil
}

func (m ST_Hint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Hint) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "default":
			*m = 1
		case "eastAsia":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Hint) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "default"
	case 2:
		return "eastAsia"
	}
	return ""
}

func (m ST_Hint) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Hint) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Theme byte

const (
	ST_ThemeUnset         ST_Theme = 0
	ST_ThemeMajorEastAsia ST_Theme = 1
	ST_ThemeMajorBidi     ST_Theme = 2
	ST_ThemeMajorAscii    ST_Theme = 3
	ST_ThemeMajorHAnsi    ST_Theme = 4
	ST_ThemeMinorEastAsia ST_Theme = 5
	ST_ThemeMinorBidi     ST_Theme = 6
	ST_ThemeMinorAscii    ST_Theme = 7
	ST_ThemeMinorHAnsi    ST_Theme = 8
)

func (e ST_Theme) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ThemeUnset:
		attr.Value = ""
	case ST_ThemeMajorEastAsia:
		attr.Value = "majorEastAsia"
	case ST_ThemeMajorBidi:
		attr.Value = "majorBidi"
	case ST_ThemeMajorAscii:
		attr.Value = "majorAscii"
	case ST_ThemeMajorHAnsi:
		attr.Value = "majorHAnsi"
	case ST_ThemeMinorEastAsia:
		attr.Value = "minorEastAsia"
	case ST_ThemeMinorBidi:
		attr.Value = "minorBidi"
	case ST_ThemeMinorAscii:
		attr.Value = "minorAscii"
	case ST_ThemeMinorHAnsi:
		attr.Value = "minorHAnsi"
	}
	return attr, nil
}

func (e *ST_Theme) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "majorEastAsia":
		*e = 1
	case "majorBidi":
		*e = 2
	case "majorAscii":
		*e = 3
	case "majorHAnsi":
		*e = 4
	case "minorEastAsia":
		*e = 5
	case "minorBidi":
		*e = 6
	case "minorAscii":
		*e = 7
	case "minorHAnsi":
		*e = 8
	}
	return nil
}

func (m ST_Theme) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Theme) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "majorEastAsia":
			*m = 1
		case "majorBidi":
			*m = 2
		case "majorAscii":
			*m = 3
		case "majorHAnsi":
			*m = 4
		case "minorEastAsia":
			*m = 5
		case "minorBidi":
			*m = 6
		case "minorAscii":
			*m = 7
		case "minorHAnsi":
			*m = 8
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Theme) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "majorEastAsia"
	case 2:
		return "majorBidi"
	case 3:
		return "majorAscii"
	case 4:
		return "majorHAnsi"
	case 5:
		return "minorEastAsia"
	case 6:
		return "minorBidi"
	case 7:
		return "minorAscii"
	case 8:
		return "minorHAnsi"
	}
	return ""
}

func (m ST_Theme) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Theme) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_RubyAlign byte

const (
	ST_RubyAlignUnset            ST_RubyAlign = 0
	ST_RubyAlignCenter           ST_RubyAlign = 1
	ST_RubyAlignDistributeLetter ST_RubyAlign = 2
	ST_RubyAlignDistributeSpace  ST_RubyAlign = 3
	ST_RubyAlignLeft             ST_RubyAlign = 4
	ST_RubyAlignRight            ST_RubyAlign = 5
	ST_RubyAlignRightVertical    ST_RubyAlign = 6
)

func (e ST_RubyAlign) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_RubyAlignUnset:
		attr.Value = ""
	case ST_RubyAlignCenter:
		attr.Value = "center"
	case ST_RubyAlignDistributeLetter:
		attr.Value = "distributeLetter"
	case ST_RubyAlignDistributeSpace:
		attr.Value = "distributeSpace"
	case ST_RubyAlignLeft:
		attr.Value = "left"
	case ST_RubyAlignRight:
		attr.Value = "right"
	case ST_RubyAlignRightVertical:
		attr.Value = "rightVertical"
	}
	return attr, nil
}

func (e *ST_RubyAlign) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "center":
		*e = 1
	case "distributeLetter":
		*e = 2
	case "distributeSpace":
		*e = 3
	case "left":
		*e = 4
	case "right":
		*e = 5
	case "rightVertical":
		*e = 6
	}
	return nil
}

func (m ST_RubyAlign) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_RubyAlign) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "center":
			*m = 1
		case "distributeLetter":
			*m = 2
		case "distributeSpace":
			*m = 3
		case "left":
			*m = 4
		case "right":
			*m = 5
		case "rightVertical":
			*m = 6
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_RubyAlign) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "center"
	case 2:
		return "distributeLetter"
	case 3:
		return "distributeSpace"
	case 4:
		return "left"
	case 5:
		return "right"
	case 6:
		return "rightVertical"
	}
	return ""
}

func (m ST_RubyAlign) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_RubyAlign) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Lock byte

const (
	ST_LockUnset            ST_Lock = 0
	ST_LockSdtLocked        ST_Lock = 1
	ST_LockContentLocked    ST_Lock = 2
	ST_LockUnlocked         ST_Lock = 3
	ST_LockSdtContentLocked ST_Lock = 4
)

func (e ST_Lock) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_LockUnset:
		attr.Value = ""
	case ST_LockSdtLocked:
		attr.Value = "sdtLocked"
	case ST_LockContentLocked:
		attr.Value = "contentLocked"
	case ST_LockUnlocked:
		attr.Value = "unlocked"
	case ST_LockSdtContentLocked:
		attr.Value = "sdtContentLocked"
	}
	return attr, nil
}

func (e *ST_Lock) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "sdtLocked":
		*e = 1
	case "contentLocked":
		*e = 2
	case "unlocked":
		*e = 3
	case "sdtContentLocked":
		*e = 4
	}
	return nil
}

func (m ST_Lock) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Lock) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "sdtLocked":
			*m = 1
		case "contentLocked":
			*m = 2
		case "unlocked":
			*m = 3
		case "sdtContentLocked":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Lock) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "sdtLocked"
	case 2:
		return "contentLocked"
	case 3:
		return "unlocked"
	case 4:
		return "sdtContentLocked"
	}
	return ""
}

func (m ST_Lock) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Lock) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_SdtDateMappingType byte

const (
	ST_SdtDateMappingTypeUnset    ST_SdtDateMappingType = 0
	ST_SdtDateMappingTypeText     ST_SdtDateMappingType = 1
	ST_SdtDateMappingTypeDate     ST_SdtDateMappingType = 2
	ST_SdtDateMappingTypeDateTime ST_SdtDateMappingType = 3
)

func (e ST_SdtDateMappingType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_SdtDateMappingTypeUnset:
		attr.Value = ""
	case ST_SdtDateMappingTypeText:
		attr.Value = "text"
	case ST_SdtDateMappingTypeDate:
		attr.Value = "date"
	case ST_SdtDateMappingTypeDateTime:
		attr.Value = "dateTime"
	}
	return attr, nil
}

func (e *ST_SdtDateMappingType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "text":
		*e = 1
	case "date":
		*e = 2
	case "dateTime":
		*e = 3
	}
	return nil
}

func (m ST_SdtDateMappingType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_SdtDateMappingType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "text":
			*m = 1
		case "date":
			*m = 2
		case "dateTime":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_SdtDateMappingType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "text"
	case 2:
		return "date"
	case 3:
		return "dateTime"
	}
	return ""
}

func (m ST_SdtDateMappingType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_SdtDateMappingType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Direction byte

const (
	ST_DirectionUnset ST_Direction = 0
	ST_DirectionLtr   ST_Direction = 1
	ST_DirectionRtl   ST_Direction = 2
)

func (e ST_Direction) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DirectionUnset:
		attr.Value = ""
	case ST_DirectionLtr:
		attr.Value = "ltr"
	case ST_DirectionRtl:
		attr.Value = "rtl"
	}
	return attr, nil
}

func (e *ST_Direction) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "ltr":
		*e = 1
	case "rtl":
		*e = 2
	}
	return nil
}

func (m ST_Direction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Direction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "ltr":
			*m = 1
		case "rtl":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Direction) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "ltr"
	case 2:
		return "rtl"
	}
	return ""
}

func (m ST_Direction) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Direction) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TblWidth byte

const (
	ST_TblWidthUnset ST_TblWidth = 0
	ST_TblWidthNil   ST_TblWidth = 1
	ST_TblWidthPct   ST_TblWidth = 2
	ST_TblWidthDxa   ST_TblWidth = 3
	ST_TblWidthAuto  ST_TblWidth = 4
)

func (e ST_TblWidth) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TblWidthUnset:
		attr.Value = ""
	case ST_TblWidthNil:
		attr.Value = "nil"
	case ST_TblWidthPct:
		attr.Value = "pct"
	case ST_TblWidthDxa:
		attr.Value = "dxa"
	case ST_TblWidthAuto:
		attr.Value = "auto"
	}
	return attr, nil
}

func (e *ST_TblWidth) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "nil":
		*e = 1
	case "pct":
		*e = 2
	case "dxa":
		*e = 3
	case "auto":
		*e = 4
	}
	return nil
}

func (m ST_TblWidth) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TblWidth) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "nil":
			*m = 1
		case "pct":
			*m = 2
		case "dxa":
			*m = 3
		case "auto":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_TblWidth) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "nil"
	case 2:
		return "pct"
	case 3:
		return "dxa"
	case 4:
		return "auto"
	}
	return ""
}

func (m ST_TblWidth) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TblWidth) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Merge byte

const (
	ST_MergeUnset    ST_Merge = 0
	ST_MergeContinue ST_Merge = 1
	ST_MergeRestart  ST_Merge = 2
)

func (e ST_Merge) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_MergeUnset:
		attr.Value = ""
	case ST_MergeContinue:
		attr.Value = "continue"
	case ST_MergeRestart:
		attr.Value = "restart"
	}
	return attr, nil
}

func (e *ST_Merge) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "continue":
		*e = 1
	case "restart":
		*e = 2
	}
	return nil
}

func (m ST_Merge) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Merge) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "continue":
			*m = 1
		case "restart":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Merge) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "continue"
	case 2:
		return "restart"
	}
	return ""
}

func (m ST_Merge) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Merge) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TblLayoutType byte

const (
	ST_TblLayoutTypeUnset   ST_TblLayoutType = 0
	ST_TblLayoutTypeFixed   ST_TblLayoutType = 1
	ST_TblLayoutTypeAutofit ST_TblLayoutType = 2
)

func (e ST_TblLayoutType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TblLayoutTypeUnset:
		attr.Value = ""
	case ST_TblLayoutTypeFixed:
		attr.Value = "fixed"
	case ST_TblLayoutTypeAutofit:
		attr.Value = "autofit"
	}
	return attr, nil
}

func (e *ST_TblLayoutType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "fixed":
		*e = 1
	case "autofit":
		*e = 2
	}
	return nil
}

func (m ST_TblLayoutType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TblLayoutType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "fixed":
			*m = 1
		case "autofit":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_TblLayoutType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "fixed"
	case 2:
		return "autofit"
	}
	return ""
}

func (m ST_TblLayoutType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TblLayoutType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TblOverlap byte

const (
	ST_TblOverlapUnset   ST_TblOverlap = 0
	ST_TblOverlapNever   ST_TblOverlap = 1
	ST_TblOverlapOverlap ST_TblOverlap = 2
)

func (e ST_TblOverlap) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TblOverlapUnset:
		attr.Value = ""
	case ST_TblOverlapNever:
		attr.Value = "never"
	case ST_TblOverlapOverlap:
		attr.Value = "overlap"
	}
	return attr, nil
}

func (e *ST_TblOverlap) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "never":
		*e = 1
	case "overlap":
		*e = 2
	}
	return nil
}

func (m ST_TblOverlap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TblOverlap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "never":
			*m = 1
		case "overlap":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_TblOverlap) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "never"
	case 2:
		return "overlap"
	}
	return ""
}

func (m ST_TblOverlap) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TblOverlap) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_FtnPos byte

const (
	ST_FtnPosUnset       ST_FtnPos = 0
	ST_FtnPosPageBottom  ST_FtnPos = 1
	ST_FtnPosBeneathText ST_FtnPos = 2
	ST_FtnPosSectEnd     ST_FtnPos = 3
	ST_FtnPosDocEnd      ST_FtnPos = 4
)

func (e ST_FtnPos) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FtnPosUnset:
		attr.Value = ""
	case ST_FtnPosPageBottom:
		attr.Value = "pageBottom"
	case ST_FtnPosBeneathText:
		attr.Value = "beneathText"
	case ST_FtnPosSectEnd:
		attr.Value = "sectEnd"
	case ST_FtnPosDocEnd:
		attr.Value = "docEnd"
	}
	return attr, nil
}

func (e *ST_FtnPos) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "pageBottom":
		*e = 1
	case "beneathText":
		*e = 2
	case "sectEnd":
		*e = 3
	case "docEnd":
		*e = 4
	}
	return nil
}

func (m ST_FtnPos) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_FtnPos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "pageBottom":
			*m = 1
		case "beneathText":
			*m = 2
		case "sectEnd":
			*m = 3
		case "docEnd":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_FtnPos) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "pageBottom"
	case 2:
		return "beneathText"
	case 3:
		return "sectEnd"
	case 4:
		return "docEnd"
	}
	return ""
}

func (m ST_FtnPos) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FtnPos) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_EdnPos byte

const (
	ST_EdnPosUnset   ST_EdnPos = 0
	ST_EdnPosSectEnd ST_EdnPos = 1
	ST_EdnPosDocEnd  ST_EdnPos = 2
)

func (e ST_EdnPos) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_EdnPosUnset:
		attr.Value = ""
	case ST_EdnPosSectEnd:
		attr.Value = "sectEnd"
	case ST_EdnPosDocEnd:
		attr.Value = "docEnd"
	}
	return attr, nil
}

func (e *ST_EdnPos) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "sectEnd":
		*e = 1
	case "docEnd":
		*e = 2
	}
	return nil
}

func (m ST_EdnPos) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_EdnPos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "sectEnd":
			*m = 1
		case "docEnd":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_EdnPos) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "sectEnd"
	case 2:
		return "docEnd"
	}
	return ""
}

func (m ST_EdnPos) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_EdnPos) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_RestartNumber byte

const (
	ST_RestartNumberUnset      ST_RestartNumber = 0
	ST_RestartNumberContinuous ST_RestartNumber = 1
	ST_RestartNumberEachSect   ST_RestartNumber = 2
	ST_RestartNumberEachPage   ST_RestartNumber = 3
)

func (e ST_RestartNumber) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_RestartNumberUnset:
		attr.Value = ""
	case ST_RestartNumberContinuous:
		attr.Value = "continuous"
	case ST_RestartNumberEachSect:
		attr.Value = "eachSect"
	case ST_RestartNumberEachPage:
		attr.Value = "eachPage"
	}
	return attr, nil
}

func (e *ST_RestartNumber) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "continuous":
		*e = 1
	case "eachSect":
		*e = 2
	case "eachPage":
		*e = 3
	}
	return nil
}

func (m ST_RestartNumber) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_RestartNumber) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "continuous":
			*m = 1
		case "eachSect":
			*m = 2
		case "eachPage":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_RestartNumber) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "continuous"
	case 2:
		return "eachSect"
	case 3:
		return "eachPage"
	}
	return ""
}

func (m ST_RestartNumber) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_RestartNumber) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_MailMergeSourceType byte

const (
	ST_MailMergeSourceTypeUnset       ST_MailMergeSourceType = 0
	ST_MailMergeSourceTypeDatabase    ST_MailMergeSourceType = 1
	ST_MailMergeSourceTypeAddressBook ST_MailMergeSourceType = 2
	ST_MailMergeSourceTypeDocument1   ST_MailMergeSourceType = 3
	ST_MailMergeSourceTypeDocument2   ST_MailMergeSourceType = 4
	ST_MailMergeSourceTypeText        ST_MailMergeSourceType = 5
	ST_MailMergeSourceTypeEmail       ST_MailMergeSourceType = 6
	ST_MailMergeSourceTypeNative      ST_MailMergeSourceType = 7
	ST_MailMergeSourceTypeLegacy      ST_MailMergeSourceType = 8
	ST_MailMergeSourceTypeMaster      ST_MailMergeSourceType = 9
)

func (e ST_MailMergeSourceType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_MailMergeSourceTypeUnset:
		attr.Value = ""
	case ST_MailMergeSourceTypeDatabase:
		attr.Value = "database"
	case ST_MailMergeSourceTypeAddressBook:
		attr.Value = "addressBook"
	case ST_MailMergeSourceTypeDocument1:
		attr.Value = "document1"
	case ST_MailMergeSourceTypeDocument2:
		attr.Value = "document2"
	case ST_MailMergeSourceTypeText:
		attr.Value = "text"
	case ST_MailMergeSourceTypeEmail:
		attr.Value = "email"
	case ST_MailMergeSourceTypeNative:
		attr.Value = "native"
	case ST_MailMergeSourceTypeLegacy:
		attr.Value = "legacy"
	case ST_MailMergeSourceTypeMaster:
		attr.Value = "master"
	}
	return attr, nil
}

func (e *ST_MailMergeSourceType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "database":
		*e = 1
	case "addressBook":
		*e = 2
	case "document1":
		*e = 3
	case "document2":
		*e = 4
	case "text":
		*e = 5
	case "email":
		*e = 6
	case "native":
		*e = 7
	case "legacy":
		*e = 8
	case "master":
		*e = 9
	}
	return nil
}

func (m ST_MailMergeSourceType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_MailMergeSourceType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "database":
			*m = 1
		case "addressBook":
			*m = 2
		case "document1":
			*m = 3
		case "document2":
			*m = 4
		case "text":
			*m = 5
		case "email":
			*m = 6
		case "native":
			*m = 7
		case "legacy":
			*m = 8
		case "master":
			*m = 9
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_MailMergeSourceType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "database"
	case 2:
		return "addressBook"
	case 3:
		return "document1"
	case 4:
		return "document2"
	case 5:
		return "text"
	case 6:
		return "email"
	case 7:
		return "native"
	case 8:
		return "legacy"
	case 9:
		return "master"
	}
	return ""
}

func (m ST_MailMergeSourceType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_MailMergeSourceType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TargetScreenSz byte

const (
	ST_TargetScreenSzUnset     ST_TargetScreenSz = 0
	ST_TargetScreenSz544x376   ST_TargetScreenSz = 1
	ST_TargetScreenSz640x480   ST_TargetScreenSz = 2
	ST_TargetScreenSz720x512   ST_TargetScreenSz = 3
	ST_TargetScreenSz800x600   ST_TargetScreenSz = 4
	ST_TargetScreenSz1024x768  ST_TargetScreenSz = 5
	ST_TargetScreenSz1152x882  ST_TargetScreenSz = 6
	ST_TargetScreenSz1152x900  ST_TargetScreenSz = 7
	ST_TargetScreenSz1280x1024 ST_TargetScreenSz = 8
	ST_TargetScreenSz1600x1200 ST_TargetScreenSz = 9
	ST_TargetScreenSz1800x1440 ST_TargetScreenSz = 10
	ST_TargetScreenSz1920x1200 ST_TargetScreenSz = 11
)

func (e ST_TargetScreenSz) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TargetScreenSzUnset:
		attr.Value = ""
	case ST_TargetScreenSz544x376:
		attr.Value = "544x376"
	case ST_TargetScreenSz640x480:
		attr.Value = "640x480"
	case ST_TargetScreenSz720x512:
		attr.Value = "720x512"
	case ST_TargetScreenSz800x600:
		attr.Value = "800x600"
	case ST_TargetScreenSz1024x768:
		attr.Value = "1024x768"
	case ST_TargetScreenSz1152x882:
		attr.Value = "1152x882"
	case ST_TargetScreenSz1152x900:
		attr.Value = "1152x900"
	case ST_TargetScreenSz1280x1024:
		attr.Value = "1280x1024"
	case ST_TargetScreenSz1600x1200:
		attr.Value = "1600x1200"
	case ST_TargetScreenSz1800x1440:
		attr.Value = "1800x1440"
	case ST_TargetScreenSz1920x1200:
		attr.Value = "1920x1200"
	}
	return attr, nil
}

func (e *ST_TargetScreenSz) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "544x376":
		*e = 1
	case "640x480":
		*e = 2
	case "720x512":
		*e = 3
	case "800x600":
		*e = 4
	case "1024x768":
		*e = 5
	case "1152x882":
		*e = 6
	case "1152x900":
		*e = 7
	case "1280x1024":
		*e = 8
	case "1600x1200":
		*e = 9
	case "1800x1440":
		*e = 10
	case "1920x1200":
		*e = 11
	}
	return nil
}

func (m ST_TargetScreenSz) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TargetScreenSz) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "544x376":
			*m = 1
		case "640x480":
			*m = 2
		case "720x512":
			*m = 3
		case "800x600":
			*m = 4
		case "1024x768":
			*m = 5
		case "1152x882":
			*m = 6
		case "1152x900":
			*m = 7
		case "1280x1024":
			*m = 8
		case "1600x1200":
			*m = 9
		case "1800x1440":
			*m = 10
		case "1920x1200":
			*m = 11
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_TargetScreenSz) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "544x376"
	case 2:
		return "640x480"
	case 3:
		return "720x512"
	case 4:
		return "800x600"
	case 5:
		return "1024x768"
	case 6:
		return "1152x882"
	case 7:
		return "1152x900"
	case 8:
		return "1280x1024"
	case 9:
		return "1600x1200"
	case 10:
		return "1800x1440"
	case 11:
		return "1920x1200"
	}
	return ""
}

func (m ST_TargetScreenSz) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TargetScreenSz) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_CharacterSpacing byte

const (
	ST_CharacterSpacingUnset                              ST_CharacterSpacing = 0
	ST_CharacterSpacingDoNotCompress                      ST_CharacterSpacing = 1
	ST_CharacterSpacingCompressPunctuation                ST_CharacterSpacing = 2
	ST_CharacterSpacingCompressPunctuationAndJapaneseKana ST_CharacterSpacing = 3
)

func (e ST_CharacterSpacing) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_CharacterSpacingUnset:
		attr.Value = ""
	case ST_CharacterSpacingDoNotCompress:
		attr.Value = "doNotCompress"
	case ST_CharacterSpacingCompressPunctuation:
		attr.Value = "compressPunctuation"
	case ST_CharacterSpacingCompressPunctuationAndJapaneseKana:
		attr.Value = "compressPunctuationAndJapaneseKana"
	}
	return attr, nil
}

func (e *ST_CharacterSpacing) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "doNotCompress":
		*e = 1
	case "compressPunctuation":
		*e = 2
	case "compressPunctuationAndJapaneseKana":
		*e = 3
	}
	return nil
}

func (m ST_CharacterSpacing) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_CharacterSpacing) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "doNotCompress":
			*m = 1
		case "compressPunctuation":
			*m = 2
		case "compressPunctuationAndJapaneseKana":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_CharacterSpacing) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "doNotCompress"
	case 2:
		return "compressPunctuation"
	case 3:
		return "compressPunctuationAndJapaneseKana"
	}
	return ""
}

func (m ST_CharacterSpacing) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_CharacterSpacing) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_WmlColorSchemeIndex byte

const (
	ST_WmlColorSchemeIndexUnset             ST_WmlColorSchemeIndex = 0
	ST_WmlColorSchemeIndexDark1             ST_WmlColorSchemeIndex = 1
	ST_WmlColorSchemeIndexLight1            ST_WmlColorSchemeIndex = 2
	ST_WmlColorSchemeIndexDark2             ST_WmlColorSchemeIndex = 3
	ST_WmlColorSchemeIndexLight2            ST_WmlColorSchemeIndex = 4
	ST_WmlColorSchemeIndexAccent1           ST_WmlColorSchemeIndex = 5
	ST_WmlColorSchemeIndexAccent2           ST_WmlColorSchemeIndex = 6
	ST_WmlColorSchemeIndexAccent3           ST_WmlColorSchemeIndex = 7
	ST_WmlColorSchemeIndexAccent4           ST_WmlColorSchemeIndex = 8
	ST_WmlColorSchemeIndexAccent5           ST_WmlColorSchemeIndex = 9
	ST_WmlColorSchemeIndexAccent6           ST_WmlColorSchemeIndex = 10
	ST_WmlColorSchemeIndexHyperlink         ST_WmlColorSchemeIndex = 11
	ST_WmlColorSchemeIndexFollowedHyperlink ST_WmlColorSchemeIndex = 12
)

func (e ST_WmlColorSchemeIndex) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_WmlColorSchemeIndexUnset:
		attr.Value = ""
	case ST_WmlColorSchemeIndexDark1:
		attr.Value = "dark1"
	case ST_WmlColorSchemeIndexLight1:
		attr.Value = "light1"
	case ST_WmlColorSchemeIndexDark2:
		attr.Value = "dark2"
	case ST_WmlColorSchemeIndexLight2:
		attr.Value = "light2"
	case ST_WmlColorSchemeIndexAccent1:
		attr.Value = "accent1"
	case ST_WmlColorSchemeIndexAccent2:
		attr.Value = "accent2"
	case ST_WmlColorSchemeIndexAccent3:
		attr.Value = "accent3"
	case ST_WmlColorSchemeIndexAccent4:
		attr.Value = "accent4"
	case ST_WmlColorSchemeIndexAccent5:
		attr.Value = "accent5"
	case ST_WmlColorSchemeIndexAccent6:
		attr.Value = "accent6"
	case ST_WmlColorSchemeIndexHyperlink:
		attr.Value = "hyperlink"
	case ST_WmlColorSchemeIndexFollowedHyperlink:
		attr.Value = "followedHyperlink"
	}
	return attr, nil
}

func (e *ST_WmlColorSchemeIndex) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "dark1":
		*e = 1
	case "light1":
		*e = 2
	case "dark2":
		*e = 3
	case "light2":
		*e = 4
	case "accent1":
		*e = 5
	case "accent2":
		*e = 6
	case "accent3":
		*e = 7
	case "accent4":
		*e = 8
	case "accent5":
		*e = 9
	case "accent6":
		*e = 10
	case "hyperlink":
		*e = 11
	case "followedHyperlink":
		*e = 12
	}
	return nil
}

func (m ST_WmlColorSchemeIndex) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_WmlColorSchemeIndex) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "dark1":
			*m = 1
		case "light1":
			*m = 2
		case "dark2":
			*m = 3
		case "light2":
			*m = 4
		case "accent1":
			*m = 5
		case "accent2":
			*m = 6
		case "accent3":
			*m = 7
		case "accent4":
			*m = 8
		case "accent5":
			*m = 9
		case "accent6":
			*m = 10
		case "hyperlink":
			*m = 11
		case "followedHyperlink":
			*m = 12
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_WmlColorSchemeIndex) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "dark1"
	case 2:
		return "light1"
	case 3:
		return "dark2"
	case 4:
		return "light2"
	case 5:
		return "accent1"
	case 6:
		return "accent2"
	case 7:
		return "accent3"
	case 8:
		return "accent4"
	case 9:
		return "accent5"
	case 10:
		return "accent6"
	case 11:
		return "hyperlink"
	case 12:
		return "followedHyperlink"
	}
	return ""
}

func (m ST_WmlColorSchemeIndex) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_WmlColorSchemeIndex) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_StyleSort byte

const (
	ST_StyleSortUnset    ST_StyleSort = 0
	ST_StyleSortName     ST_StyleSort = 1
	ST_StyleSortPriority ST_StyleSort = 2
	ST_StyleSortDefault  ST_StyleSort = 3
	ST_StyleSortFont     ST_StyleSort = 4
	ST_StyleSortBasedOn  ST_StyleSort = 5
	ST_StyleSortType     ST_StyleSort = 6
	ST_StyleSort0000     ST_StyleSort = 7
	ST_StyleSort0001     ST_StyleSort = 8
	ST_StyleSort0002     ST_StyleSort = 9
	ST_StyleSort0003     ST_StyleSort = 10
	ST_StyleSort0004     ST_StyleSort = 11
	ST_StyleSort0005     ST_StyleSort = 12
)

func (e ST_StyleSort) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_StyleSortUnset:
		attr.Value = ""
	case ST_StyleSortName:
		attr.Value = "name"
	case ST_StyleSortPriority:
		attr.Value = "priority"
	case ST_StyleSortDefault:
		attr.Value = "default"
	case ST_StyleSortFont:
		attr.Value = "font"
	case ST_StyleSortBasedOn:
		attr.Value = "basedOn"
	case ST_StyleSortType:
		attr.Value = "type"
	case ST_StyleSort0000:
		attr.Value = "0000"
	case ST_StyleSort0001:
		attr.Value = "0001"
	case ST_StyleSort0002:
		attr.Value = "0002"
	case ST_StyleSort0003:
		attr.Value = "0003"
	case ST_StyleSort0004:
		attr.Value = "0004"
	case ST_StyleSort0005:
		attr.Value = "0005"
	}
	return attr, nil
}

func (e *ST_StyleSort) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "name":
		*e = 1
	case "priority":
		*e = 2
	case "default":
		*e = 3
	case "font":
		*e = 4
	case "basedOn":
		*e = 5
	case "type":
		*e = 6
	case "0000":
		*e = 7
	case "0001":
		*e = 8
	case "0002":
		*e = 9
	case "0003":
		*e = 10
	case "0004":
		*e = 11
	case "0005":
		*e = 12
	}
	return nil
}

func (m ST_StyleSort) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_StyleSort) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "name":
			*m = 1
		case "priority":
			*m = 2
		case "default":
			*m = 3
		case "font":
			*m = 4
		case "basedOn":
			*m = 5
		case "type":
			*m = 6
		case "0000":
			*m = 7
		case "0001":
			*m = 8
		case "0002":
			*m = 9
		case "0003":
			*m = 10
		case "0004":
			*m = 11
		case "0005":
			*m = 12
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_StyleSort) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "name"
	case 2:
		return "priority"
	case 3:
		return "default"
	case 4:
		return "font"
	case 5:
		return "basedOn"
	case 6:
		return "type"
	case 7:
		return "0000"
	case 8:
		return "0001"
	case 9:
		return "0002"
	case 10:
		return "0003"
	case 11:
		return "0004"
	case 12:
		return "0005"
	}
	return ""
}

func (m ST_StyleSort) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_StyleSort) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_FrameScrollbar byte

const (
	ST_FrameScrollbarUnset ST_FrameScrollbar = 0
	ST_FrameScrollbarOn    ST_FrameScrollbar = 1
	ST_FrameScrollbarOff   ST_FrameScrollbar = 2
	ST_FrameScrollbarAuto  ST_FrameScrollbar = 3
)

func (e ST_FrameScrollbar) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FrameScrollbarUnset:
		attr.Value = ""
	case ST_FrameScrollbarOn:
		attr.Value = "on"
	case ST_FrameScrollbarOff:
		attr.Value = "off"
	case ST_FrameScrollbarAuto:
		attr.Value = "auto"
	}
	return attr, nil
}

func (e *ST_FrameScrollbar) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "on":
		*e = 1
	case "off":
		*e = 2
	case "auto":
		*e = 3
	}
	return nil
}

func (m ST_FrameScrollbar) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_FrameScrollbar) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "on":
			*m = 1
		case "off":
			*m = 2
		case "auto":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_FrameScrollbar) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "on"
	case 2:
		return "off"
	case 3:
		return "auto"
	}
	return ""
}

func (m ST_FrameScrollbar) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FrameScrollbar) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_FrameLayout byte

const (
	ST_FrameLayoutUnset ST_FrameLayout = 0
	ST_FrameLayoutRows  ST_FrameLayout = 1
	ST_FrameLayoutCols  ST_FrameLayout = 2
	ST_FrameLayoutNone  ST_FrameLayout = 3
)

func (e ST_FrameLayout) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FrameLayoutUnset:
		attr.Value = ""
	case ST_FrameLayoutRows:
		attr.Value = "rows"
	case ST_FrameLayoutCols:
		attr.Value = "cols"
	case ST_FrameLayoutNone:
		attr.Value = "none"
	}
	return attr, nil
}

func (e *ST_FrameLayout) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "rows":
		*e = 1
	case "cols":
		*e = 2
	case "none":
		*e = 3
	}
	return nil
}

func (m ST_FrameLayout) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_FrameLayout) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "rows":
			*m = 1
		case "cols":
			*m = 2
		case "none":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_FrameLayout) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "rows"
	case 2:
		return "cols"
	case 3:
		return "none"
	}
	return ""
}

func (m ST_FrameLayout) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FrameLayout) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_LevelSuffix byte

const (
	ST_LevelSuffixUnset   ST_LevelSuffix = 0
	ST_LevelSuffixTab     ST_LevelSuffix = 1
	ST_LevelSuffixSpace   ST_LevelSuffix = 2
	ST_LevelSuffixNothing ST_LevelSuffix = 3
)

func (e ST_LevelSuffix) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_LevelSuffixUnset:
		attr.Value = ""
	case ST_LevelSuffixTab:
		attr.Value = "tab"
	case ST_LevelSuffixSpace:
		attr.Value = "space"
	case ST_LevelSuffixNothing:
		attr.Value = "nothing"
	}
	return attr, nil
}

func (e *ST_LevelSuffix) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "tab":
		*e = 1
	case "space":
		*e = 2
	case "nothing":
		*e = 3
	}
	return nil
}

func (m ST_LevelSuffix) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_LevelSuffix) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "tab":
			*m = 1
		case "space":
			*m = 2
		case "nothing":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_LevelSuffix) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "tab"
	case 2:
		return "space"
	case 3:
		return "nothing"
	}
	return ""
}

func (m ST_LevelSuffix) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_LevelSuffix) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_MultiLevelType byte

const (
	ST_MultiLevelTypeUnset            ST_MultiLevelType = 0
	ST_MultiLevelTypeSingleLevel      ST_MultiLevelType = 1
	ST_MultiLevelTypeMultilevel       ST_MultiLevelType = 2
	ST_MultiLevelTypeHybridMultilevel ST_MultiLevelType = 3
)

func (e ST_MultiLevelType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_MultiLevelTypeUnset:
		attr.Value = ""
	case ST_MultiLevelTypeSingleLevel:
		attr.Value = "singleLevel"
	case ST_MultiLevelTypeMultilevel:
		attr.Value = "multilevel"
	case ST_MultiLevelTypeHybridMultilevel:
		attr.Value = "hybridMultilevel"
	}
	return attr, nil
}

func (e *ST_MultiLevelType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "singleLevel":
		*e = 1
	case "multilevel":
		*e = 2
	case "hybridMultilevel":
		*e = 3
	}
	return nil
}

func (m ST_MultiLevelType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_MultiLevelType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "singleLevel":
			*m = 1
		case "multilevel":
			*m = 2
		case "hybridMultilevel":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_MultiLevelType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "singleLevel"
	case 2:
		return "multilevel"
	case 3:
		return "hybridMultilevel"
	}
	return ""
}

func (m ST_MultiLevelType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_MultiLevelType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TblStyleOverrideType byte

const (
	ST_TblStyleOverrideTypeUnset      ST_TblStyleOverrideType = 0
	ST_TblStyleOverrideTypeWholeTable ST_TblStyleOverrideType = 1
	ST_TblStyleOverrideTypeFirstRow   ST_TblStyleOverrideType = 2
	ST_TblStyleOverrideTypeLastRow    ST_TblStyleOverrideType = 3
	ST_TblStyleOverrideTypeFirstCol   ST_TblStyleOverrideType = 4
	ST_TblStyleOverrideTypeLastCol    ST_TblStyleOverrideType = 5
	ST_TblStyleOverrideTypeBand1Vert  ST_TblStyleOverrideType = 6
	ST_TblStyleOverrideTypeBand2Vert  ST_TblStyleOverrideType = 7
	ST_TblStyleOverrideTypeBand1Horz  ST_TblStyleOverrideType = 8
	ST_TblStyleOverrideTypeBand2Horz  ST_TblStyleOverrideType = 9
	ST_TblStyleOverrideTypeNeCell     ST_TblStyleOverrideType = 10
	ST_TblStyleOverrideTypeNwCell     ST_TblStyleOverrideType = 11
	ST_TblStyleOverrideTypeSeCell     ST_TblStyleOverrideType = 12
	ST_TblStyleOverrideTypeSwCell     ST_TblStyleOverrideType = 13
)

func (e ST_TblStyleOverrideType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TblStyleOverrideTypeUnset:
		attr.Value = ""
	case ST_TblStyleOverrideTypeWholeTable:
		attr.Value = "wholeTable"
	case ST_TblStyleOverrideTypeFirstRow:
		attr.Value = "firstRow"
	case ST_TblStyleOverrideTypeLastRow:
		attr.Value = "lastRow"
	case ST_TblStyleOverrideTypeFirstCol:
		attr.Value = "firstCol"
	case ST_TblStyleOverrideTypeLastCol:
		attr.Value = "lastCol"
	case ST_TblStyleOverrideTypeBand1Vert:
		attr.Value = "band1Vert"
	case ST_TblStyleOverrideTypeBand2Vert:
		attr.Value = "band2Vert"
	case ST_TblStyleOverrideTypeBand1Horz:
		attr.Value = "band1Horz"
	case ST_TblStyleOverrideTypeBand2Horz:
		attr.Value = "band2Horz"
	case ST_TblStyleOverrideTypeNeCell:
		attr.Value = "neCell"
	case ST_TblStyleOverrideTypeNwCell:
		attr.Value = "nwCell"
	case ST_TblStyleOverrideTypeSeCell:
		attr.Value = "seCell"
	case ST_TblStyleOverrideTypeSwCell:
		attr.Value = "swCell"
	}
	return attr, nil
}

func (e *ST_TblStyleOverrideType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "wholeTable":
		*e = 1
	case "firstRow":
		*e = 2
	case "lastRow":
		*e = 3
	case "firstCol":
		*e = 4
	case "lastCol":
		*e = 5
	case "band1Vert":
		*e = 6
	case "band2Vert":
		*e = 7
	case "band1Horz":
		*e = 8
	case "band2Horz":
		*e = 9
	case "neCell":
		*e = 10
	case "nwCell":
		*e = 11
	case "seCell":
		*e = 12
	case "swCell":
		*e = 13
	}
	return nil
}

func (m ST_TblStyleOverrideType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TblStyleOverrideType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "wholeTable":
			*m = 1
		case "firstRow":
			*m = 2
		case "lastRow":
			*m = 3
		case "firstCol":
			*m = 4
		case "lastCol":
			*m = 5
		case "band1Vert":
			*m = 6
		case "band2Vert":
			*m = 7
		case "band1Horz":
			*m = 8
		case "band2Horz":
			*m = 9
		case "neCell":
			*m = 10
		case "nwCell":
			*m = 11
		case "seCell":
			*m = 12
		case "swCell":
			*m = 13
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_TblStyleOverrideType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "wholeTable"
	case 2:
		return "firstRow"
	case 3:
		return "lastRow"
	case 4:
		return "firstCol"
	case 5:
		return "lastCol"
	case 6:
		return "band1Vert"
	case 7:
		return "band2Vert"
	case 8:
		return "band1Horz"
	case 9:
		return "band2Horz"
	case 10:
		return "neCell"
	case 11:
		return "nwCell"
	case 12:
		return "seCell"
	case 13:
		return "swCell"
	}
	return ""
}

func (m ST_TblStyleOverrideType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TblStyleOverrideType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_StyleType byte

const (
	ST_StyleTypeUnset     ST_StyleType = 0
	ST_StyleTypeParagraph ST_StyleType = 1
	ST_StyleTypeCharacter ST_StyleType = 2
	ST_StyleTypeTable     ST_StyleType = 3
	ST_StyleTypeNumbering ST_StyleType = 4
)

func (e ST_StyleType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_StyleTypeUnset:
		attr.Value = ""
	case ST_StyleTypeParagraph:
		attr.Value = "paragraph"
	case ST_StyleTypeCharacter:
		attr.Value = "character"
	case ST_StyleTypeTable:
		attr.Value = "table"
	case ST_StyleTypeNumbering:
		attr.Value = "numbering"
	}
	return attr, nil
}

func (e *ST_StyleType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "paragraph":
		*e = 1
	case "character":
		*e = 2
	case "table":
		*e = 3
	case "numbering":
		*e = 4
	}
	return nil
}

func (m ST_StyleType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_StyleType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "paragraph":
			*m = 1
		case "character":
			*m = 2
		case "table":
			*m = 3
		case "numbering":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_StyleType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "paragraph"
	case 2:
		return "character"
	case 3:
		return "table"
	case 4:
		return "numbering"
	}
	return ""
}

func (m ST_StyleType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_StyleType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_FontFamily byte

const (
	ST_FontFamilyUnset      ST_FontFamily = 0
	ST_FontFamilyDecorative ST_FontFamily = 1
	ST_FontFamilyModern     ST_FontFamily = 2
	ST_FontFamilyRoman      ST_FontFamily = 3
	ST_FontFamilyScript     ST_FontFamily = 4
	ST_FontFamilySwiss      ST_FontFamily = 5
	ST_FontFamilyAuto       ST_FontFamily = 6
)

func (e ST_FontFamily) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FontFamilyUnset:
		attr.Value = ""
	case ST_FontFamilyDecorative:
		attr.Value = "decorative"
	case ST_FontFamilyModern:
		attr.Value = "modern"
	case ST_FontFamilyRoman:
		attr.Value = "roman"
	case ST_FontFamilyScript:
		attr.Value = "script"
	case ST_FontFamilySwiss:
		attr.Value = "swiss"
	case ST_FontFamilyAuto:
		attr.Value = "auto"
	}
	return attr, nil
}

func (e *ST_FontFamily) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "decorative":
		*e = 1
	case "modern":
		*e = 2
	case "roman":
		*e = 3
	case "script":
		*e = 4
	case "swiss":
		*e = 5
	case "auto":
		*e = 6
	}
	return nil
}

func (m ST_FontFamily) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_FontFamily) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "decorative":
			*m = 1
		case "modern":
			*m = 2
		case "roman":
			*m = 3
		case "script":
			*m = 4
		case "swiss":
			*m = 5
		case "auto":
			*m = 6
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_FontFamily) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "decorative"
	case 2:
		return "modern"
	case 3:
		return "roman"
	case 4:
		return "script"
	case 5:
		return "swiss"
	case 6:
		return "auto"
	}
	return ""
}

func (m ST_FontFamily) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FontFamily) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Pitch byte

const (
	ST_PitchUnset    ST_Pitch = 0
	ST_PitchFixed    ST_Pitch = 1
	ST_PitchVariable ST_Pitch = 2
	ST_PitchDefault  ST_Pitch = 3
)

func (e ST_Pitch) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PitchUnset:
		attr.Value = ""
	case ST_PitchFixed:
		attr.Value = "fixed"
	case ST_PitchVariable:
		attr.Value = "variable"
	case ST_PitchDefault:
		attr.Value = "default"
	}
	return attr, nil
}

func (e *ST_Pitch) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "fixed":
		*e = 1
	case "variable":
		*e = 2
	case "default":
		*e = 3
	}
	return nil
}

func (m ST_Pitch) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Pitch) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "fixed":
			*m = 1
		case "variable":
			*m = 2
		case "default":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Pitch) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "fixed"
	case 2:
		return "variable"
	case 3:
		return "default"
	}
	return ""
}

func (m ST_Pitch) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Pitch) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ThemeColor byte

const (
	ST_ThemeColorUnset             ST_ThemeColor = 0
	ST_ThemeColorDark1             ST_ThemeColor = 1
	ST_ThemeColorLight1            ST_ThemeColor = 2
	ST_ThemeColorDark2             ST_ThemeColor = 3
	ST_ThemeColorLight2            ST_ThemeColor = 4
	ST_ThemeColorAccent1           ST_ThemeColor = 5
	ST_ThemeColorAccent2           ST_ThemeColor = 6
	ST_ThemeColorAccent3           ST_ThemeColor = 7
	ST_ThemeColorAccent4           ST_ThemeColor = 8
	ST_ThemeColorAccent5           ST_ThemeColor = 9
	ST_ThemeColorAccent6           ST_ThemeColor = 10
	ST_ThemeColorHyperlink         ST_ThemeColor = 11
	ST_ThemeColorFollowedHyperlink ST_ThemeColor = 12
	ST_ThemeColorNone              ST_ThemeColor = 13
	ST_ThemeColorBackground1       ST_ThemeColor = 14
	ST_ThemeColorText1             ST_ThemeColor = 15
	ST_ThemeColorBackground2       ST_ThemeColor = 16
	ST_ThemeColorText2             ST_ThemeColor = 17
)

func (e ST_ThemeColor) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ThemeColorUnset:
		attr.Value = ""
	case ST_ThemeColorDark1:
		attr.Value = "dark1"
	case ST_ThemeColorLight1:
		attr.Value = "light1"
	case ST_ThemeColorDark2:
		attr.Value = "dark2"
	case ST_ThemeColorLight2:
		attr.Value = "light2"
	case ST_ThemeColorAccent1:
		attr.Value = "accent1"
	case ST_ThemeColorAccent2:
		attr.Value = "accent2"
	case ST_ThemeColorAccent3:
		attr.Value = "accent3"
	case ST_ThemeColorAccent4:
		attr.Value = "accent4"
	case ST_ThemeColorAccent5:
		attr.Value = "accent5"
	case ST_ThemeColorAccent6:
		attr.Value = "accent6"
	case ST_ThemeColorHyperlink:
		attr.Value = "hyperlink"
	case ST_ThemeColorFollowedHyperlink:
		attr.Value = "followedHyperlink"
	case ST_ThemeColorNone:
		attr.Value = "none"
	case ST_ThemeColorBackground1:
		attr.Value = "background1"
	case ST_ThemeColorText1:
		attr.Value = "text1"
	case ST_ThemeColorBackground2:
		attr.Value = "background2"
	case ST_ThemeColorText2:
		attr.Value = "text2"
	}
	return attr, nil
}

func (e *ST_ThemeColor) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "dark1":
		*e = 1
	case "light1":
		*e = 2
	case "dark2":
		*e = 3
	case "light2":
		*e = 4
	case "accent1":
		*e = 5
	case "accent2":
		*e = 6
	case "accent3":
		*e = 7
	case "accent4":
		*e = 8
	case "accent5":
		*e = 9
	case "accent6":
		*e = 10
	case "hyperlink":
		*e = 11
	case "followedHyperlink":
		*e = 12
	case "none":
		*e = 13
	case "background1":
		*e = 14
	case "text1":
		*e = 15
	case "background2":
		*e = 16
	case "text2":
		*e = 17
	}
	return nil
}

func (m ST_ThemeColor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ThemeColor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "dark1":
			*m = 1
		case "light1":
			*m = 2
		case "dark2":
			*m = 3
		case "light2":
			*m = 4
		case "accent1":
			*m = 5
		case "accent2":
			*m = 6
		case "accent3":
			*m = 7
		case "accent4":
			*m = 8
		case "accent5":
			*m = 9
		case "accent6":
			*m = 10
		case "hyperlink":
			*m = 11
		case "followedHyperlink":
			*m = 12
		case "none":
			*m = 13
		case "background1":
			*m = 14
		case "text1":
			*m = 15
		case "background2":
			*m = 16
		case "text2":
			*m = 17
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ThemeColor) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "dark1"
	case 2:
		return "light1"
	case 3:
		return "dark2"
	case 4:
		return "light2"
	case 5:
		return "accent1"
	case 6:
		return "accent2"
	case 7:
		return "accent3"
	case 8:
		return "accent4"
	case 9:
		return "accent5"
	case 10:
		return "accent6"
	case 11:
		return "hyperlink"
	case 12:
		return "followedHyperlink"
	case 13:
		return "none"
	case 14:
		return "background1"
	case 15:
		return "text1"
	case 16:
		return "background2"
	case 17:
		return "text2"
	}
	return ""
}

func (m ST_ThemeColor) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ThemeColor) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DocPartBehavior byte

const (
	ST_DocPartBehaviorUnset   ST_DocPartBehavior = 0
	ST_DocPartBehaviorContent ST_DocPartBehavior = 1
	ST_DocPartBehaviorP       ST_DocPartBehavior = 2
	ST_DocPartBehaviorPg      ST_DocPartBehavior = 3
)

func (e ST_DocPartBehavior) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DocPartBehaviorUnset:
		attr.Value = ""
	case ST_DocPartBehaviorContent:
		attr.Value = "content"
	case ST_DocPartBehaviorP:
		attr.Value = "p"
	case ST_DocPartBehaviorPg:
		attr.Value = "pg"
	}
	return attr, nil
}

func (e *ST_DocPartBehavior) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "content":
		*e = 1
	case "p":
		*e = 2
	case "pg":
		*e = 3
	}
	return nil
}

func (m ST_DocPartBehavior) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_DocPartBehavior) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "content":
			*m = 1
		case "p":
			*m = 2
		case "pg":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_DocPartBehavior) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "content"
	case 2:
		return "p"
	case 3:
		return "pg"
	}
	return ""
}

func (m ST_DocPartBehavior) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_DocPartBehavior) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DocPartType byte

const (
	ST_DocPartTypeUnset    ST_DocPartType = 0
	ST_DocPartTypeNone     ST_DocPartType = 1
	ST_DocPartTypeNormal   ST_DocPartType = 2
	ST_DocPartTypeAutoExp  ST_DocPartType = 3
	ST_DocPartTypeToolbar  ST_DocPartType = 4
	ST_DocPartTypeSpeller  ST_DocPartType = 5
	ST_DocPartTypeFormFld  ST_DocPartType = 6
	ST_DocPartTypeBbPlcHdr ST_DocPartType = 7
)

func (e ST_DocPartType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DocPartTypeUnset:
		attr.Value = ""
	case ST_DocPartTypeNone:
		attr.Value = "none"
	case ST_DocPartTypeNormal:
		attr.Value = "normal"
	case ST_DocPartTypeAutoExp:
		attr.Value = "autoExp"
	case ST_DocPartTypeToolbar:
		attr.Value = "toolbar"
	case ST_DocPartTypeSpeller:
		attr.Value = "speller"
	case ST_DocPartTypeFormFld:
		attr.Value = "formFld"
	case ST_DocPartTypeBbPlcHdr:
		attr.Value = "bbPlcHdr"
	}
	return attr, nil
}

func (e *ST_DocPartType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "normal":
		*e = 2
	case "autoExp":
		*e = 3
	case "toolbar":
		*e = 4
	case "speller":
		*e = 5
	case "formFld":
		*e = 6
	case "bbPlcHdr":
		*e = 7
	}
	return nil
}

func (m ST_DocPartType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_DocPartType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "normal":
			*m = 2
		case "autoExp":
			*m = 3
		case "toolbar":
			*m = 4
		case "speller":
			*m = 5
		case "formFld":
			*m = 6
		case "bbPlcHdr":
			*m = 7
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_DocPartType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "normal"
	case 3:
		return "autoExp"
	case 4:
		return "toolbar"
	case 5:
		return "speller"
	case 6:
		return "formFld"
	case 7:
		return "bbPlcHdr"
	}
	return ""
}

func (m ST_DocPartType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_DocPartType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DocPartGallery byte

const (
	ST_DocPartGalleryUnset             ST_DocPartGallery = 0
	ST_DocPartGalleryPlaceholder       ST_DocPartGallery = 1
	ST_DocPartGalleryAny               ST_DocPartGallery = 2
	ST_DocPartGalleryDefault           ST_DocPartGallery = 3
	ST_DocPartGalleryDocParts          ST_DocPartGallery = 4
	ST_DocPartGalleryCoverPg           ST_DocPartGallery = 5
	ST_DocPartGalleryEq                ST_DocPartGallery = 6
	ST_DocPartGalleryFtrs              ST_DocPartGallery = 7
	ST_DocPartGalleryHdrs              ST_DocPartGallery = 8
	ST_DocPartGalleryPgNum             ST_DocPartGallery = 9
	ST_DocPartGalleryTbls              ST_DocPartGallery = 10
	ST_DocPartGalleryWatermarks        ST_DocPartGallery = 11
	ST_DocPartGalleryAutoTxt           ST_DocPartGallery = 12
	ST_DocPartGalleryTxtBox            ST_DocPartGallery = 13
	ST_DocPartGalleryPgNumT            ST_DocPartGallery = 14
	ST_DocPartGalleryPgNumB            ST_DocPartGallery = 15
	ST_DocPartGalleryPgNumMargins      ST_DocPartGallery = 16
	ST_DocPartGalleryTblOfContents     ST_DocPartGallery = 17
	ST_DocPartGalleryBib               ST_DocPartGallery = 18
	ST_DocPartGalleryCustQuickParts    ST_DocPartGallery = 19
	ST_DocPartGalleryCustCoverPg       ST_DocPartGallery = 20
	ST_DocPartGalleryCustEq            ST_DocPartGallery = 21
	ST_DocPartGalleryCustFtrs          ST_DocPartGallery = 22
	ST_DocPartGalleryCustHdrs          ST_DocPartGallery = 23
	ST_DocPartGalleryCustPgNum         ST_DocPartGallery = 24
	ST_DocPartGalleryCustTbls          ST_DocPartGallery = 25
	ST_DocPartGalleryCustWatermarks    ST_DocPartGallery = 26
	ST_DocPartGalleryCustAutoTxt       ST_DocPartGallery = 27
	ST_DocPartGalleryCustTxtBox        ST_DocPartGallery = 28
	ST_DocPartGalleryCustPgNumT        ST_DocPartGallery = 29
	ST_DocPartGalleryCustPgNumB        ST_DocPartGallery = 30
	ST_DocPartGalleryCustPgNumMargins  ST_DocPartGallery = 31
	ST_DocPartGalleryCustTblOfContents ST_DocPartGallery = 32
	ST_DocPartGalleryCustBib           ST_DocPartGallery = 33
	ST_DocPartGalleryCustom1           ST_DocPartGallery = 34
	ST_DocPartGalleryCustom2           ST_DocPartGallery = 35
	ST_DocPartGalleryCustom3           ST_DocPartGallery = 36
	ST_DocPartGalleryCustom4           ST_DocPartGallery = 37
	ST_DocPartGalleryCustom5           ST_DocPartGallery = 38
)

func (e ST_DocPartGallery) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DocPartGalleryUnset:
		attr.Value = ""
	case ST_DocPartGalleryPlaceholder:
		attr.Value = "placeholder"
	case ST_DocPartGalleryAny:
		attr.Value = "any"
	case ST_DocPartGalleryDefault:
		attr.Value = "default"
	case ST_DocPartGalleryDocParts:
		attr.Value = "docParts"
	case ST_DocPartGalleryCoverPg:
		attr.Value = "coverPg"
	case ST_DocPartGalleryEq:
		attr.Value = "eq"
	case ST_DocPartGalleryFtrs:
		attr.Value = "ftrs"
	case ST_DocPartGalleryHdrs:
		attr.Value = "hdrs"
	case ST_DocPartGalleryPgNum:
		attr.Value = "pgNum"
	case ST_DocPartGalleryTbls:
		attr.Value = "tbls"
	case ST_DocPartGalleryWatermarks:
		attr.Value = "watermarks"
	case ST_DocPartGalleryAutoTxt:
		attr.Value = "autoTxt"
	case ST_DocPartGalleryTxtBox:
		attr.Value = "txtBox"
	case ST_DocPartGalleryPgNumT:
		attr.Value = "pgNumT"
	case ST_DocPartGalleryPgNumB:
		attr.Value = "pgNumB"
	case ST_DocPartGalleryPgNumMargins:
		attr.Value = "pgNumMargins"
	case ST_DocPartGalleryTblOfContents:
		attr.Value = "tblOfContents"
	case ST_DocPartGalleryBib:
		attr.Value = "bib"
	case ST_DocPartGalleryCustQuickParts:
		attr.Value = "custQuickParts"
	case ST_DocPartGalleryCustCoverPg:
		attr.Value = "custCoverPg"
	case ST_DocPartGalleryCustEq:
		attr.Value = "custEq"
	case ST_DocPartGalleryCustFtrs:
		attr.Value = "custFtrs"
	case ST_DocPartGalleryCustHdrs:
		attr.Value = "custHdrs"
	case ST_DocPartGalleryCustPgNum:
		attr.Value = "custPgNum"
	case ST_DocPartGalleryCustTbls:
		attr.Value = "custTbls"
	case ST_DocPartGalleryCustWatermarks:
		attr.Value = "custWatermarks"
	case ST_DocPartGalleryCustAutoTxt:
		attr.Value = "custAutoTxt"
	case ST_DocPartGalleryCustTxtBox:
		attr.Value = "custTxtBox"
	case ST_DocPartGalleryCustPgNumT:
		attr.Value = "custPgNumT"
	case ST_DocPartGalleryCustPgNumB:
		attr.Value = "custPgNumB"
	case ST_DocPartGalleryCustPgNumMargins:
		attr.Value = "custPgNumMargins"
	case ST_DocPartGalleryCustTblOfContents:
		attr.Value = "custTblOfContents"
	case ST_DocPartGalleryCustBib:
		attr.Value = "custBib"
	case ST_DocPartGalleryCustom1:
		attr.Value = "custom1"
	case ST_DocPartGalleryCustom2:
		attr.Value = "custom2"
	case ST_DocPartGalleryCustom3:
		attr.Value = "custom3"
	case ST_DocPartGalleryCustom4:
		attr.Value = "custom4"
	case ST_DocPartGalleryCustom5:
		attr.Value = "custom5"
	}
	return attr, nil
}

func (e *ST_DocPartGallery) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "placeholder":
		*e = 1
	case "any":
		*e = 2
	case "default":
		*e = 3
	case "docParts":
		*e = 4
	case "coverPg":
		*e = 5
	case "eq":
		*e = 6
	case "ftrs":
		*e = 7
	case "hdrs":
		*e = 8
	case "pgNum":
		*e = 9
	case "tbls":
		*e = 10
	case "watermarks":
		*e = 11
	case "autoTxt":
		*e = 12
	case "txtBox":
		*e = 13
	case "pgNumT":
		*e = 14
	case "pgNumB":
		*e = 15
	case "pgNumMargins":
		*e = 16
	case "tblOfContents":
		*e = 17
	case "bib":
		*e = 18
	case "custQuickParts":
		*e = 19
	case "custCoverPg":
		*e = 20
	case "custEq":
		*e = 21
	case "custFtrs":
		*e = 22
	case "custHdrs":
		*e = 23
	case "custPgNum":
		*e = 24
	case "custTbls":
		*e = 25
	case "custWatermarks":
		*e = 26
	case "custAutoTxt":
		*e = 27
	case "custTxtBox":
		*e = 28
	case "custPgNumT":
		*e = 29
	case "custPgNumB":
		*e = 30
	case "custPgNumMargins":
		*e = 31
	case "custTblOfContents":
		*e = 32
	case "custBib":
		*e = 33
	case "custom1":
		*e = 34
	case "custom2":
		*e = 35
	case "custom3":
		*e = 36
	case "custom4":
		*e = 37
	case "custom5":
		*e = 38
	}
	return nil
}

func (m ST_DocPartGallery) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_DocPartGallery) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "placeholder":
			*m = 1
		case "any":
			*m = 2
		case "default":
			*m = 3
		case "docParts":
			*m = 4
		case "coverPg":
			*m = 5
		case "eq":
			*m = 6
		case "ftrs":
			*m = 7
		case "hdrs":
			*m = 8
		case "pgNum":
			*m = 9
		case "tbls":
			*m = 10
		case "watermarks":
			*m = 11
		case "autoTxt":
			*m = 12
		case "txtBox":
			*m = 13
		case "pgNumT":
			*m = 14
		case "pgNumB":
			*m = 15
		case "pgNumMargins":
			*m = 16
		case "tblOfContents":
			*m = 17
		case "bib":
			*m = 18
		case "custQuickParts":
			*m = 19
		case "custCoverPg":
			*m = 20
		case "custEq":
			*m = 21
		case "custFtrs":
			*m = 22
		case "custHdrs":
			*m = 23
		case "custPgNum":
			*m = 24
		case "custTbls":
			*m = 25
		case "custWatermarks":
			*m = 26
		case "custAutoTxt":
			*m = 27
		case "custTxtBox":
			*m = 28
		case "custPgNumT":
			*m = 29
		case "custPgNumB":
			*m = 30
		case "custPgNumMargins":
			*m = 31
		case "custTblOfContents":
			*m = 32
		case "custBib":
			*m = 33
		case "custom1":
			*m = 34
		case "custom2":
			*m = 35
		case "custom3":
			*m = 36
		case "custom4":
			*m = 37
		case "custom5":
			*m = 38
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_DocPartGallery) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "placeholder"
	case 2:
		return "any"
	case 3:
		return "default"
	case 4:
		return "docParts"
	case 5:
		return "coverPg"
	case 6:
		return "eq"
	case 7:
		return "ftrs"
	case 8:
		return "hdrs"
	case 9:
		return "pgNum"
	case 10:
		return "tbls"
	case 11:
		return "watermarks"
	case 12:
		return "autoTxt"
	case 13:
		return "txtBox"
	case 14:
		return "pgNumT"
	case 15:
		return "pgNumB"
	case 16:
		return "pgNumMargins"
	case 17:
		return "tblOfContents"
	case 18:
		return "bib"
	case 19:
		return "custQuickParts"
	case 20:
		return "custCoverPg"
	case 21:
		return "custEq"
	case 22:
		return "custFtrs"
	case 23:
		return "custHdrs"
	case 24:
		return "custPgNum"
	case 25:
		return "custTbls"
	case 26:
		return "custWatermarks"
	case 27:
		return "custAutoTxt"
	case 28:
		return "custTxtBox"
	case 29:
		return "custPgNumT"
	case 30:
		return "custPgNumB"
	case 31:
		return "custPgNumMargins"
	case 32:
		return "custTblOfContents"
	case 33:
		return "custBib"
	case 34:
		return "custom1"
	case 35:
		return "custom2"
	case 36:
		return "custom3"
	case 37:
		return "custom4"
	case 38:
		return "custom5"
	}
	return ""
}

func (m ST_DocPartGallery) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_DocPartGallery) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_CaptionPos byte

const (
	ST_CaptionPosUnset ST_CaptionPos = 0
	ST_CaptionPosAbove ST_CaptionPos = 1
	ST_CaptionPosBelow ST_CaptionPos = 2
	ST_CaptionPosLeft  ST_CaptionPos = 3
	ST_CaptionPosRight ST_CaptionPos = 4
)

func (e ST_CaptionPos) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_CaptionPosUnset:
		attr.Value = ""
	case ST_CaptionPosAbove:
		attr.Value = "above"
	case ST_CaptionPosBelow:
		attr.Value = "below"
	case ST_CaptionPosLeft:
		attr.Value = "left"
	case ST_CaptionPosRight:
		attr.Value = "right"
	}
	return attr, nil
}

func (e *ST_CaptionPos) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "above":
		*e = 1
	case "below":
		*e = 2
	case "left":
		*e = 3
	case "right":
		*e = 4
	}
	return nil
}

func (m ST_CaptionPos) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_CaptionPos) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "above":
			*m = 1
		case "below":
			*m = 2
		case "left":
			*m = 3
		case "right":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_CaptionPos) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "above"
	case 2:
		return "below"
	case 3:
		return "left"
	case 4:
		return "right"
	}
	return ""
}

func (m ST_CaptionPos) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_CaptionPos) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Empty", NewCT_Empty)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_OnOff", NewCT_OnOff)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_LongHexNumber", NewCT_LongHexNumber)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Charset", NewCT_Charset)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DecimalNumber", NewCT_DecimalNumber)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_UnsignedDecimalNumber", NewCT_UnsignedDecimalNumber)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DecimalNumberOrPrecent", NewCT_DecimalNumberOrPrecent)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TwipsMeasure", NewCT_TwipsMeasure)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SignedTwipsMeasure", NewCT_SignedTwipsMeasure)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_PixelsMeasure", NewCT_PixelsMeasure)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_HpsMeasure", NewCT_HpsMeasure)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SignedHpsMeasure", NewCT_SignedHpsMeasure)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_MacroName", NewCT_MacroName)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_String", NewCT_String)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TextScale", NewCT_TextScale)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Highlight", NewCT_Highlight)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Color", NewCT_Color)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Lang", NewCT_Lang)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Guid", NewCT_Guid)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Underline", NewCT_Underline)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TextEffect", NewCT_TextEffect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Border", NewCT_Border)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Shd", NewCT_Shd)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_VerticalAlignRun", NewCT_VerticalAlignRun)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FitText", NewCT_FitText)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Em", NewCT_Em)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Language", NewCT_Language)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_EastAsianLayout", NewCT_EastAsianLayout)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FramePr", NewCT_FramePr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TabStop", NewCT_TabStop)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Spacing", NewCT_Spacing)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Ind", NewCT_Ind)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Jc", NewCT_Jc)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_JcTable", NewCT_JcTable)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_View", NewCT_View)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Zoom", NewCT_Zoom)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_WritingStyle", NewCT_WritingStyle)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Proof", NewCT_Proof)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DocType", NewCT_DocType)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DocProtect", NewCT_DocProtect)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_MailMergeDocType", NewCT_MailMergeDocType)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_MailMergeDataType", NewCT_MailMergeDataType)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_MailMergeDest", NewCT_MailMergeDest)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_MailMergeOdsoFMDFieldType", NewCT_MailMergeOdsoFMDFieldType)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TrackChangesView", NewCT_TrackChangesView)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Kinsoku", NewCT_Kinsoku)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TextDirection", NewCT_TextDirection)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TextAlignment", NewCT_TextAlignment)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Markup", NewCT_Markup)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TrackChange", NewCT_TrackChange)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_CellMergeTrackChange", NewCT_CellMergeTrackChange)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TrackChangeRange", NewCT_TrackChangeRange)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_MarkupRange", NewCT_MarkupRange)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_BookmarkRange", NewCT_BookmarkRange)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Bookmark", NewCT_Bookmark)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_MoveBookmark", NewCT_MoveBookmark)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Comment", NewCT_Comment)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TrackChangeNumbering", NewCT_TrackChangeNumbering)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TblPrExChange", NewCT_TblPrExChange)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TcPrChange", NewCT_TcPrChange)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TrPrChange", NewCT_TrPrChange)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TblGridChange", NewCT_TblGridChange)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TblPrChange", NewCT_TblPrChange)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SectPrChange", NewCT_SectPrChange)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_PPrChange", NewCT_PPrChange)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_RPrChange", NewCT_RPrChange)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_ParaRPrChange", NewCT_ParaRPrChange)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_RunTrackChange", NewCT_RunTrackChange)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_NumPr", NewCT_NumPr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_PBdr", NewCT_PBdr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Tabs", NewCT_Tabs)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TextboxTightWrap", NewCT_TextboxTightWrap)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_PPr", NewCT_PPr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_PPrBase", NewCT_PPrBase)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_PPrGeneral", NewCT_PPrGeneral)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Control", NewCT_Control)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Background", NewCT_Background)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Rel", NewCT_Rel)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Object", NewCT_Object)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Picture", NewCT_Picture)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_ObjectEmbed", NewCT_ObjectEmbed)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_ObjectLink", NewCT_ObjectLink)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Drawing", NewCT_Drawing)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SimpleField", NewCT_SimpleField)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FFTextType", NewCT_FFTextType)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FFName", NewCT_FFName)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FldChar", NewCT_FldChar)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Hyperlink", NewCT_Hyperlink)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FFData", NewCT_FFData)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FFHelpText", NewCT_FFHelpText)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FFStatusText", NewCT_FFStatusText)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FFCheckBox", NewCT_FFCheckBox)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FFDDList", NewCT_FFDDList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FFTextInput", NewCT_FFTextInput)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SectType", NewCT_SectType)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_PaperSource", NewCT_PaperSource)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_PageSz", NewCT_PageSz)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_PageMar", NewCT_PageMar)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_PageBorders", NewCT_PageBorders)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_PageBorder", NewCT_PageBorder)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_BottomPageBorder", NewCT_BottomPageBorder)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TopPageBorder", NewCT_TopPageBorder)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_LineNumber", NewCT_LineNumber)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_PageNumber", NewCT_PageNumber)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Column", NewCT_Column)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Columns", NewCT_Columns)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_VerticalJc", NewCT_VerticalJc)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DocGrid", NewCT_DocGrid)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_HdrFtrRef", NewCT_HdrFtrRef)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_HdrFtr", NewCT_HdrFtr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SectPrBase", NewCT_SectPrBase)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SectPr", NewCT_SectPr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Br", NewCT_Br)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_PTab", NewCT_PTab)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Sym", NewCT_Sym)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_ProofErr", NewCT_ProofErr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Perm", NewCT_Perm)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_PermStart", NewCT_PermStart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Text", NewCT_Text)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_R", NewCT_R)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Fonts", NewCT_Fonts)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_RPr", NewCT_RPr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_MathCtrlIns", NewCT_MathCtrlIns)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_MathCtrlDel", NewCT_MathCtrlDel)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_RPrOriginal", NewCT_RPrOriginal)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_ParaRPrOriginal", NewCT_ParaRPrOriginal)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_ParaRPr", NewCT_ParaRPr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_AltChunk", NewCT_AltChunk)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_AltChunkPr", NewCT_AltChunkPr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_RubyAlign", NewCT_RubyAlign)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_RubyPr", NewCT_RubyPr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_RubyContent", NewCT_RubyContent)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Ruby", NewCT_Ruby)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Lock", NewCT_Lock)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SdtListItem", NewCT_SdtListItem)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SdtDateMappingType", NewCT_SdtDateMappingType)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_CalendarType", NewCT_CalendarType)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SdtDate", NewCT_SdtDate)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SdtComboBox", NewCT_SdtComboBox)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SdtDocPart", NewCT_SdtDocPart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SdtDropDownList", NewCT_SdtDropDownList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Placeholder", NewCT_Placeholder)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SdtText", NewCT_SdtText)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DataBinding", NewCT_DataBinding)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SdtPr", NewCT_SdtPr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SdtEndPr", NewCT_SdtEndPr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DirContentRun", NewCT_DirContentRun)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_BdoContentRun", NewCT_BdoContentRun)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SdtContentRun", NewCT_SdtContentRun)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SdtContentBlock", NewCT_SdtContentBlock)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SdtContentRow", NewCT_SdtContentRow)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SdtContentCell", NewCT_SdtContentCell)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SdtBlock", NewCT_SdtBlock)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SdtRun", NewCT_SdtRun)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SdtCell", NewCT_SdtCell)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SdtRow", NewCT_SdtRow)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Attr", NewCT_Attr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_CustomXmlRun", NewCT_CustomXmlRun)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SmartTagRun", NewCT_SmartTagRun)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_CustomXmlBlock", NewCT_CustomXmlBlock)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_CustomXmlPr", NewCT_CustomXmlPr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_CustomXmlRow", NewCT_CustomXmlRow)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_CustomXmlCell", NewCT_CustomXmlCell)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SmartTagPr", NewCT_SmartTagPr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_P", NewCT_P)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Height", NewCT_Height)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TblWidth", NewCT_TblWidth)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TblGridCol", NewCT_TblGridCol)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TblGridBase", NewCT_TblGridBase)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TblGrid", NewCT_TblGrid)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TcBorders", NewCT_TcBorders)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TcMar", NewCT_TcMar)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_VMerge", NewCT_VMerge)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_HMerge", NewCT_HMerge)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TcPrBase", NewCT_TcPrBase)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TcPr", NewCT_TcPr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TcPrInner", NewCT_TcPrInner)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Tc", NewCT_Tc)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Cnf", NewCT_Cnf)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Headers", NewCT_Headers)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TrPrBase", NewCT_TrPrBase)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TrPr", NewCT_TrPr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Row", NewCT_Row)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TblLayoutType", NewCT_TblLayoutType)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TblOverlap", NewCT_TblOverlap)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TblPPr", NewCT_TblPPr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TblCellMar", NewCT_TblCellMar)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TblBorders", NewCT_TblBorders)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TblPrBase", NewCT_TblPrBase)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TblPr", NewCT_TblPr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TblPrExBase", NewCT_TblPrExBase)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TblPrEx", NewCT_TblPrEx)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Tbl", NewCT_Tbl)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TblLook", NewCT_TblLook)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FtnPos", NewCT_FtnPos)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_EdnPos", NewCT_EdnPos)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_NumFmt", NewCT_NumFmt)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_NumRestart", NewCT_NumRestart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FtnEdnRef", NewCT_FtnEdnRef)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FtnEdnSepRef", NewCT_FtnEdnSepRef)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FtnEdn", NewCT_FtnEdn)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FtnProps", NewCT_FtnProps)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_EdnProps", NewCT_EdnProps)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FtnDocProps", NewCT_FtnDocProps)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_EdnDocProps", NewCT_EdnDocProps)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_RecipientData", NewCT_RecipientData)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Base64Binary", NewCT_Base64Binary)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Recipients", NewCT_Recipients)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_OdsoFieldMapData", NewCT_OdsoFieldMapData)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_MailMergeSourceType", NewCT_MailMergeSourceType)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Odso", NewCT_Odso)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_MailMerge", NewCT_MailMerge)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TargetScreenSz", NewCT_TargetScreenSz)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Compat", NewCT_Compat)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_CompatSetting", NewCT_CompatSetting)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DocVar", NewCT_DocVar)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DocVars", NewCT_DocVars)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DocRsids", NewCT_DocRsids)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_CharacterSpacing", NewCT_CharacterSpacing)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SaveThroughXslt", NewCT_SaveThroughXslt)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_RPrDefault", NewCT_RPrDefault)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_PPrDefault", NewCT_PPrDefault)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DocDefaults", NewCT_DocDefaults)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_ColorSchemeMapping", NewCT_ColorSchemeMapping)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_ReadingModeInkLockDown", NewCT_ReadingModeInkLockDown)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_WriteProtection", NewCT_WriteProtection)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Settings", NewCT_Settings)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_StyleSort", NewCT_StyleSort)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_StylePaneFilter", NewCT_StylePaneFilter)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_WebSettings", NewCT_WebSettings)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FrameScrollbar", NewCT_FrameScrollbar)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_OptimizeForBrowser", NewCT_OptimizeForBrowser)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Frame", NewCT_Frame)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FrameLayout", NewCT_FrameLayout)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FramesetSplitbar", NewCT_FramesetSplitbar)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Frameset", NewCT_Frameset)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_NumPicBullet", NewCT_NumPicBullet)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_LevelSuffix", NewCT_LevelSuffix)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_LevelText", NewCT_LevelText)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_LvlLegacy", NewCT_LvlLegacy)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Lvl", NewCT_Lvl)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_MultiLevelType", NewCT_MultiLevelType)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_AbstractNum", NewCT_AbstractNum)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_NumLvl", NewCT_NumLvl)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Num", NewCT_Num)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Numbering", NewCT_Numbering)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TblStylePr", NewCT_TblStylePr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Style", NewCT_Style)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_LsdException", NewCT_LsdException)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_LatentStyles", NewCT_LatentStyles)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Styles", NewCT_Styles)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Panose", NewCT_Panose)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FontFamily", NewCT_FontFamily)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Pitch", NewCT_Pitch)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FontSig", NewCT_FontSig)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FontRel", NewCT_FontRel)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Font", NewCT_Font)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_FontsList", NewCT_FontsList)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DivBdr", NewCT_DivBdr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Div", NewCT_Div)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Divs", NewCT_Divs)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_TxbxContent", NewCT_TxbxContent)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Body", NewCT_Body)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_ShapeDefaults", NewCT_ShapeDefaults)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Comments", NewCT_Comments)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Footnotes", NewCT_Footnotes)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Endnotes", NewCT_Endnotes)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_SmartTagType", NewCT_SmartTagType)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DocPartBehavior", NewCT_DocPartBehavior)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DocPartBehaviors", NewCT_DocPartBehaviors)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DocPartType", NewCT_DocPartType)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DocPartTypes", NewCT_DocPartTypes)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DocPartGallery", NewCT_DocPartGallery)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DocPartCategory", NewCT_DocPartCategory)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DocPartName", NewCT_DocPartName)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DocPartPr", NewCT_DocPartPr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DocPart", NewCT_DocPart)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DocParts", NewCT_DocParts)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Caption", NewCT_Caption)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_AutoCaption", NewCT_AutoCaption)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_AutoCaptions", NewCT_AutoCaptions)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Captions", NewCT_Captions)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_DocumentBase", NewCT_DocumentBase)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_Document", NewCT_Document)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "CT_GlossaryDocument", NewCT_GlossaryDocument)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "recipients", NewRecipients)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "txbxContent", NewTxbxContent)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "comments", NewComments)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "footnotes", NewFootnotes)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "endnotes", NewEndnotes)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "hdr", NewHdr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "ftr", NewFtr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "settings", NewSettings)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "webSettings", NewWebSettings)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "fonts", NewFonts)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "numbering", NewNumbering)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "styles", NewStyles)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "document", NewDocument)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "glossaryDocument", NewGlossaryDocument)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_PContentMath", NewEG_PContentMath)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_PContentBase", NewEG_PContentBase)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_ContentRunContentBase", NewEG_ContentRunContentBase)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_CellMarkupElements", NewEG_CellMarkupElements)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_RangeMarkupElements", NewEG_RangeMarkupElements)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_HdrFtrReferences", NewEG_HdrFtrReferences)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_SectPrContents", NewEG_SectPrContents)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_RunInnerContent", NewEG_RunInnerContent)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_RPrBase", NewEG_RPrBase)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_RPrContent", NewEG_RPrContent)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_RPr", NewEG_RPr)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_RPrMath", NewEG_RPrMath)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_ParaRPrTrackChanges", NewEG_ParaRPrTrackChanges)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_RubyContent", NewEG_RubyContent)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_ContentRunContent", NewEG_ContentRunContent)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_ContentBlockContent", NewEG_ContentBlockContent)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_ContentRowContent", NewEG_ContentRowContent)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_ContentCellContent", NewEG_ContentCellContent)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_PContent", NewEG_PContent)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_FtnEdnNumProps", NewEG_FtnEdnNumProps)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_MathContent", NewEG_MathContent)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_BlockLevelChunkElts", NewEG_BlockLevelChunkElts)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_BlockLevelElts", NewEG_BlockLevelElts)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "EG_RunLevelElts", NewEG_RunLevelElts)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "AG_Password", NewAG_Password)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "AG_TransitionalPassword", NewAG_TransitionalPassword)
	gooxml.RegisterConstructor("http://schemas.openxmlformats.org/wordprocessingml/2006/main", "AG_SectPrAttributes", NewAG_SectPrAttributes)
}
