package detailedTestCaseExecutionsUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func (detailedTestCaseExecutionsUIObject *DetailedTestCaseExecutionsUIModelStruct) CreateDetailedTestCaseExecutionsTabPage() (
	detailedTestCaseExecutionsTabPage *fyne.Container) {

	/*
		var colorDefinitionLabel *widget.Label
		colorDefinitionLabel = &widget.Label{
			BaseWidget: widget.BaseWidget{},
			Text:       "Execution Status color definition",
			Alignment:  0,
			Wrapping:   0,
			TextStyle: fyne.TextStyle{
				Bold:      true,
				Italic:    false,
				Monospace: false,
				Symbol:    false,
				TabWidth:  0,
			},
		}
	*/

	detailedTestCaseExecutionsTabPage = container.New(layout.NewVBoxLayout(), layout.NewSpacer(), detailedTestCaseExecutionsUIObject.generateExecutionColorPalette())

	return detailedTestCaseExecutionsTabPage

}

func (detailedTestCaseExecutionsUIObject *DetailedTestCaseExecutionsUIModelStruct) generateExecutionColorPalette() (
	executionColorPaletteContainerObject *fyne.Container) {

	/*
		#b6d7a8	INITIATED = 0; // All set up for execution, but has not been triggered to start execution
		#ffff00	EXECUTING = 1; // TestInstruction is execution
		#4a86e8	CONTROLLED_INTERRUPTION = 2; // Interrupted by in a controlled way
		#4a86e8	CONTROLLED_INTERRUPTION_CAN_BE_RERUN = 3; // Interrupted by in a controlled way, but can be rerun
		#00ff00	FINISHED_OK = 4; // Finish as expected to TestInstruction definition
		#00ff00	FINISHED_OK_CAN_BE_RERUN = 5; // Finish as expected to TestInstruction definition, but can be rerun
		#ff0000	FINISHED_NOT_OK = 6; // Finish with errors in validations
		#ff0000	FINISHED_NOT_OK_CAN_BE_RERUN = 7; // Finish with errors in validations, but can be rerun
		#9900ff	UNEXPECTED_INTERRUPTION = 8; // The TestInstruction stopped executed in an unexpected way
		#9900ff	UNEXPECTED_INTERRUPTION_CAN_BE_RERUN = 9; // The TestInstruction stopped executed in an unexpected way, but can be rerun
		#fbbc04	TIMEOUT_INTERRUPTION = 10; // The TestInstruction had a forced stop because of timeout due to {time.Now() > 'ExpectedExecutionEndTimeStamp'}
		#fbbc04	TIMEOUT_INTERRUPTION_CAN_BE_RERUN = 11; // The TestInstruction had a forced stop because of timeout due to {time.Now() > 'ExpectedExecutionEndTimeStamp'}, but can be rerun
	*/

	// Define Header
	var colorDefinitionLabel *widget.Label
	colorDefinitionLabel = &widget.Label{
		BaseWidget: widget.BaseWidget{},
		Text:       "Execution Status color definition",
		Alignment:  fyne.TextAlignLeading,
		Wrapping:   0,
		TextStyle: fyne.TextStyle{
			Bold:      true,
			Italic:    false,
			Monospace: false,
			Symbol:    false,
			TabWidth:  0,
		},
	}

	// Execution status 'INITIATED = 0;'
	//initiatedLabel := widget.NewLabel("Initiated")
	initiatedText := canvas.Text{
		Alignment: fyne.TextAlignCenter,
		Color: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF,
		},
		Text:     " Initiated ",
		TextSize: 15,
		TextStyle: fyne.TextStyle{
			Bold:      false,
			Italic:    false,
			Monospace: false,
			Symbol:    false,
			TabWidth:  0,
		},
	}
	initiatedTextbox := canvas.NewRectangle(color.RGBA{
		R: 0xb6,
		G: 0xd7,
		B: 0xa8,
		A: 0xFF,
	})
	initiatedContainer := container.New(layout.NewMaxLayout(), initiatedTextbox, &initiatedText)

	// Execution status 'EXECUTING = 1'
	//executingLabel := widget.NewLabel("Executing")
	executingText := canvas.Text{
		Alignment: fyne.TextAlignCenter,
		Color: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF,
		},
		Text:     " Executing ",
		TextSize: 15,
		TextStyle: fyne.TextStyle{
			Bold:      false,
			Italic:    false,
			Monospace: false,
			Symbol:    false,
			TabWidth:  0,
		},
	}
	executingTextbox := canvas.NewRectangle(color.RGBA{
		R: 0xff,
		G: 0xff,
		B: 0x00,
		A: 0xFF,
	})
	executingContainer := container.New(layout.NewMaxLayout(), executingTextbox, &executingText)

	// Execution status 'CONTROLLED_INTERRUPTION = 2'
	//controlledInterruptionLabel := widget.NewLabel("Controlled Interruption")
	controlledInterruptionText := canvas.Text{
		Alignment: fyne.TextAlignCenter,
		Color: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF,
		},
		Text:     " Controlled Interruption ",
		TextSize: 15,
		TextStyle: fyne.TextStyle{
			Bold:      false,
			Italic:    false,
			Monospace: false,
			Symbol:    false,
			TabWidth:  0,
		},
	}
	controlledInterruptionTextbox := canvas.NewRectangle(color.RGBA{
		R: 0x4a,
		G: 0x86,
		B: 0xe8,
		A: 0xFF,
	})
	controlledInterruptionContainer := container.New(layout.NewMaxLayout(), controlledInterruptionTextbox, &controlledInterruptionText)

	// Execution status 'CONTROLLED_INTERRUPTION_CAN_BE_RERUN = 3'
	//controlledInterruptionRerunLabel := widget.NewLabel("Controlled Interruption - Can Be Rerun")
	controlledInterruptionRerunText := canvas.Text{
		Alignment: fyne.TextAlignCenter,
		Color: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF,
		},
		Text:     " Controlled Interruption - Can Be Rerun ",
		TextSize: 15,
		TextStyle: fyne.TextStyle{
			Bold:      false,
			Italic:    false,
			Monospace: false,
			Symbol:    false,
			TabWidth:  0,
		},
	}
	controlledInterruptionRerunTextbox := canvas.NewRectangle(color.RGBA{
		R: 0x4a,
		G: 0x86,
		B: 0xe8,
		A: 0xFF,
	})
	controlledInterruptionRerunTextbox.StrokeColor = color.RGBA{
		R: 0x00,
		G: 0xFF,
		B: 0x00,
		A: 0xFF,
	}
	controlledInterruptionRerunTextbox.StrokeWidth = 4

	controlledInterruptionRerunContainer := container.New(layout.NewMaxLayout(), controlledInterruptionRerunTextbox, &controlledInterruptionRerunText)

	// Execution status 'FINISHED_OK = 4'
	//FinishedOkLabel := widget.NewLabel("Finished OK")
	FinishedOkText := canvas.Text{
		Alignment: fyne.TextAlignCenter,
		Color: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF,
		},
		Text:     " Finished OK ",
		TextSize: 15,
		TextStyle: fyne.TextStyle{
			Bold:      false,
			Italic:    false,
			Monospace: false,
			Symbol:    false,
			TabWidth:  0,
		},
	}
	FinishedOkTextbox := canvas.NewRectangle(color.RGBA{
		R: 0x00,
		G: 0xff,
		B: 0x00,
		A: 0xFF,
	})
	FinishedOkContainer := container.New(layout.NewMaxLayout(), FinishedOkTextbox, &FinishedOkText)

	// Execution status 'FINISHED_NOT_OK = 6'
	//finishedNotOkLabel := widget.NewLabel("Finished Not OK")
	finishedNotOkText := canvas.Text{
		Alignment: fyne.TextAlignCenter,
		Color: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF,
		},
		Text:     " Finished Not OK ",
		TextSize: 15,
		TextStyle: fyne.TextStyle{
			Bold:      false,
			Italic:    false,
			Monospace: false,
			Symbol:    false,
			TabWidth:  0,
		},
	}
	finishedNotOkTextbox := canvas.NewRectangle(color.RGBA{
		R: 0xff,
		G: 0x00,
		B: 0x00,
		A: 0xFF,
	})
	finishedNotOkContainer := container.New(layout.NewMaxLayout(), finishedNotOkTextbox, &finishedNotOkText)

	// Execution status 'FINISHED_NOT_OK_CAN_BE_RERUN = 7'
	//finishedNotOkRerunLabel := widget.NewLabel("Finished Not OK - Can Be Rerun")
	finishedNotOkRerunText := canvas.Text{
		Alignment: fyne.TextAlignCenter,
		Color: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF,
		},
		Text:     " Finished Not OK - Can Be Rerun ",
		TextSize: 15,
		TextStyle: fyne.TextStyle{
			Bold:      false,
			Italic:    false,
			Monospace: false,
			Symbol:    false,
			TabWidth:  0,
		},
	}
	finishedNotOkRerunTextbox := canvas.NewRectangle(color.RGBA{
		R: 0xff,
		G: 0x00,
		B: 0x00,
		A: 0xFF,
	})
	finishedNotOkRerunTextbox.StrokeColor = color.RGBA{
		R: 0x00,
		G: 0xFF,
		B: 0x00,
		A: 0xFF,
	}
	finishedNotOkRerunTextbox.StrokeWidth = 4

	finishedNotOkRerunContainer := container.New(layout.NewMaxLayout(), finishedNotOkRerunTextbox, &finishedNotOkRerunText)

	// Execution status 'UNEXPECTED_INTERRUPTION = 8'
	//unexpectedInterruptionLabel := widget.NewLabel("Unexpected Interruption")
	unexpectedInterruptionText := canvas.Text{
		Alignment: fyne.TextAlignCenter,
		Color: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF,
		},
		Text:     " Unexpected Interruption ",
		TextSize: 15,
		TextStyle: fyne.TextStyle{
			Bold:      false,
			Italic:    false,
			Monospace: false,
			Symbol:    false,
			TabWidth:  0,
		},
	}
	unexpectedInterruptionTextbox := canvas.NewRectangle(color.RGBA{
		R: 0x99,
		G: 0x00,
		B: 0xff,
		A: 0xFF,
	})
	unexpectedInterruptionContainer := container.New(layout.NewMaxLayout(), unexpectedInterruptionTextbox, &unexpectedInterruptionText)

	// Execution status 'UNEXPECTED_INTERRUPTION_CAN_BE_RERUN = 9'
	//unexpectedInterruptionRerunLabel := widget.NewLabel("Unexpected Interruption - Can Be Rerun")
	unexpectedInterruptionRerunText := canvas.Text{
		Alignment: fyne.TextAlignCenter,
		Color: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF,
		},
		Text:     " Unexpected Interruption - Can Be Rerun ",
		TextSize: 15,
		TextStyle: fyne.TextStyle{
			Bold:      false,
			Italic:    false,
			Monospace: false,
			Symbol:    false,
			TabWidth:  0,
		},
	}
	unexpectedInterruptionRerunTextbox := canvas.NewRectangle(color.RGBA{
		R: 0x99,
		G: 0x00,
		B: 0xff,
		A: 0xFF,
	})
	unexpectedInterruptionRerunTextbox.StrokeColor = color.RGBA{
		R: 0x00,
		G: 0xFF,
		B: 0x00,
		A: 0xFF,
	}
	unexpectedInterruptionRerunTextbox.StrokeWidth = 4

	unexpectedInterruptionRerunContainer := container.New(layout.NewMaxLayout(), unexpectedInterruptionRerunTextbox, &unexpectedInterruptionRerunText)

	// Execution status 'TIMEOUT_INTERRUPTION = 10'
	//timeoutInterruptionLabel := widget.NewLabel("Timeout Interruption")
	timeoutInterruptionText := canvas.Text{
		Alignment: fyne.TextAlignCenter,
		Color: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF,
		},
		Text:     " Timeout Interruption ",
		TextSize: 15,
		TextStyle: fyne.TextStyle{
			Bold:      false,
			Italic:    false,
			Monospace: false,
			Symbol:    false,
			TabWidth:  0,
		},
	}
	timeoutInterruptionTextbox := canvas.NewRectangle(color.RGBA{
		R: 0xfb,
		G: 0xbc,
		B: 0x04,
		A: 0xFF,
	})
	timeoutInterruptionContainer := container.New(layout.NewMaxLayout(), timeoutInterruptionTextbox, &timeoutInterruptionText)

	// Execution status 'TIMEOUT_INTERRUPTION_CAN_BE_RERUN = 11'
	//timeoutInterruptionRerunLabel := widget.NewLabel("Timeout Interruption - Can Be Rerun")
	timeoutInterruptionRerunText := canvas.Text{
		Alignment: fyne.TextAlignCenter,
		Color: color.RGBA{
			R: 0x00,
			G: 0x00,
			B: 0x00,
			A: 0xFF,
		},
		Text:     " Timeout Interruption - Can Be Rerun ",
		TextSize: 15,
		TextStyle: fyne.TextStyle{
			Bold:      false,
			Italic:    false,
			Monospace: false,
			Symbol:    false,
			TabWidth:  0,
		},
	}
	timeoutInterruptionRerunTextbox := canvas.NewRectangle(color.RGBA{
		R: 0xfb,
		G: 0xbc,
		B: 0x04,
		A: 0xFF,
	})
	timeoutInterruptionRerunTextbox.StrokeColor = color.RGBA{
		R: 0x00,
		G: 0xFF,
		B: 0x00,
		A: 0xFF,
	}
	timeoutInterruptionRerunTextbox.StrokeWidth = 4

	timeoutInterruptionRerunContainer := container.New(layout.NewMaxLayout(), timeoutInterruptionRerunTextbox, &timeoutInterruptionRerunText)

	// Create the color palette
	var executionColorPalette *fyne.Container
	executionColorPalette = container.New(
		layout.NewHBoxLayout(),
		initiatedContainer,
		executingContainer,
		controlledInterruptionContainer,
		controlledInterruptionRerunContainer,
		FinishedOkContainer,
		finishedNotOkContainer,
		finishedNotOkRerunContainer,
		unexpectedInterruptionContainer,
		unexpectedInterruptionRerunContainer,
		timeoutInterruptionContainer,
		timeoutInterruptionRerunContainer)

	// Create full Color Palett canvas object
	executionColorPaletteContainerObject = container.New(layout.NewVBoxLayout(), colorDefinitionLabel, executionColorPalette)

	return executionColorPaletteContainerObject
}
