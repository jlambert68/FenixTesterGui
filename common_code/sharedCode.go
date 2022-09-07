package sharedCode

import (
	"errors"
	"fmt"
	"image/color"
	"strconv"
)

// ConvertRGBAHexStringIntoRGBAColor- Converts a colors in a hex-string into 'color.RGBA'-format. "#FF03AFFF"
func ConvertRGBAHexStringIntoRGBAColor(rgbaHexString string) (rgbaValue color.RGBA, err error) {

	var (
		extractedColorRed          uint8
		extractedColorGreen        uint8
		extractedColorBlue         uint8
		extractedAlphaColorChannel uint8
	)

	errorColor := color.RGBA{
		R: 0xFF,
		G: 0x00,
		B: 0xFF,
		A: 0xFF}

	// Checka that ther String is of correct length, '#FFEEBB33'
	if len(rgbaHexString) != 9 {
		errorId := "93789d03-f728-40da-a6bd-78f8a96628a5"
		err = errors.New(fmt.Sprintf("color string with hexvalues, '%s', has not the correct lenght, '#AABBCCDDEE' in testcase with sourceUuid '%s' [ErrorID: %s]", rgbaHexString, errorId))

		return errorColor, err
	}

	hexRed := rgbaHexString[1:3]
	hexGreen := rgbaHexString[3:5]
	hexBlue := rgbaHexString[5:7]
	hexAlpha := rgbaHexString[7:9]

	// Hex-conversion for Red
	valueRed, hexConvertionErr := strconv.ParseInt(hexRed, 16, 64)
	if hexConvertionErr != nil {
		// Hex conversion failed
		errorId := "162667e9-d35e-45be-b1b4-d07877a3cd2c"
		err = errors.New(fmt.Sprintf("hexConvertion for Color failed with error message: '%s' [ErrorID: %s]", hexConvertionErr, errorId))

		return errorColor, err

	}

	// Hex-conversion for Green
	valueGreen, hexConvertionErr := strconv.ParseInt(hexGreen, 16, 64)
	if hexConvertionErr != nil {
		// Hex conversion failed
		errorId := "b2b9fae0-30e3-49df-99d7-5662b78311a3"
		err = errors.New(fmt.Sprintf("hexConvertion for Color failed with error message: '%s' [ErrorID: %s]", hexConvertionErr, errorId))

		return errorColor, err

	}

	// Hex-conversion for Blue
	valueBlue, hexConvertionErr := strconv.ParseInt(hexBlue, 16, 64)
	if hexConvertionErr != nil {
		// Hex conversion failed
		errorId := "b2b9fae0-30e3-49df-99d7-5662b78311a3"
		err = errors.New(fmt.Sprintf("hexConvertion for Color failed with error message: '%s' [ErrorID: %s]", hexConvertionErr, errorId))

		return errorColor, err

	}

	// Hex-conversion for Alpha
	valueAlpha, hexConvertionErr := strconv.ParseInt(hexAlpha, 16, 64)
	if hexConvertionErr != nil {
		// Hex conversion failed
		errorId := "f5569252-41f5-48db-8a0a-b217b1460f7d"
		err = errors.New(fmt.Sprintf("hexConvertion for Color failed with error message: '%s' [ErrorID: %s]", hexConvertionErr, errorId))

		return errorColor, err

	}

	// Convert to color values
	extractedColorRed = uint8(valueRed)
	extractedColorGreen = uint8(valueGreen)
	extractedColorBlue = uint8(valueBlue)
	extractedAlphaColorChannel = uint8(valueAlpha)

	rgbaValue = color.RGBA{
		R: extractedColorRed,
		G: extractedColorGreen,
		B: extractedColorBlue,
		A: extractedAlphaColorChannel,
	}

	return rgbaValue, err

}
