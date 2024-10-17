package soundEngine

import (
	sharedCode "FenixTesterGui/common_code"
	"bytes"
	_ "embed"
	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
	"github.com/sirupsen/logrus"
	"io"
)

// Embed System Notification
//
//go:embed soundfiles/system-notification-4-206493.mp3
var systemNotificationAsByteArray []byte

// Embed Invalid Selection
//
//go:embed soundfiles/invalid-selection-39351.mp3
var invalidSelectionAsByteArray []byte

type SoundType uint

const soundChannelSize uint16 = 100
const (
	SystemNotificationSound SoundType = iota
	InvalidNotificationSound
)

var PlaySoundChannel chan SoundType

// Remember that you should **not** create more than one context
var otoCtx *oto.Context

// Player for 'System Notification'
var systemNotificationPlayer *oto.Player
var invalidNotificationPlayer *oto.Player

// The reader for the Sound player
func playerChannelReader() {

	var soundToPlay SoundType

	// Wait for sound to be received on the channel
	soundToPlay = <-PlaySoundChannel

	switch soundToPlay {

	case SystemNotificationSound:
		go playSystemNotification()

	case InvalidNotificationSound:
		go playInvalidNotification()

	default:
		sharedCode.Logger.WithFields(logrus.Fields{
			"ID":          "12e0a633-45d4-45fd-be8b-69459cdd75c2",
			"soundToPlay": soundToPlay,
		}).Fatal("An unhandled sound was received on 'PlaySoundChannel'")

	}

}

// Initiate PlayerChannelEngine
func initiatePlayerChannelEngine() {
	PlaySoundChannel = make(chan SoundType, soundChannelSize)

	go playerChannelReader()

}

// Init the Sound Enigine if that hasn't been done
func initSoundEngine() {
	if otoCtx != nil {
		return
	}

	// Initiate PlayerChannelEngine
	initiatePlayerChannelEngine()

	// Initiate Sound Engine to be able to play all sounds, from memory
	initiateSystemNotification()
	initiateInvalidNotification()

}

func CloseDownSoundEngine() {

	var err error

	// If you don't want the player/sound anymore simply close
	if systemNotificationPlayer != nil {
		err = systemNotificationPlayer.Close()
		if err != nil {
			panic("player.Close failed: " + err.Error())
		}
	}

	if invalidNotificationPlayer != nil {
		err = invalidNotificationPlayer.Close()
		if err != nil {
			panic("player.Close failed: " + err.Error())
		}
	}

}

func initiateSystemNotification() {
	// Read the mp3 file into memory
	//fileBytes, err := os.ReadFile("./system-notification-4-206493.mp3")
	//if err != nil {
	//	panic("reading my-file.mp3 failed: " + err.Error())
	//}

	var err error

	// Convert the pure bytes into a reader object that can be used with the mp3 decoder
	var fileBytesReader *bytes.Reader
	fileBytesReader = bytes.NewReader(systemNotificationAsByteArray)

	// Decode file
	var decodedMp3 *mp3.Decoder
	decodedMp3, err = mp3.NewDecoder(fileBytesReader)
	if err != nil {
		panic("mp3.NewDecoder failed: " + err.Error())
	}

	// Prepare an Oto context (this will use your default audio device) that will
	// play all our sounds. Its configuration can't be changed later.

	var op *oto.NewContextOptions
	op = &oto.NewContextOptions{}

	// Usually 44100 or 48000. Other values might cause distortions in Oto
	op.SampleRate = 44100

	// Number of channels (aka locations) to play sounds from. Either 1 or 2.
	// 1 is mono sound, and 2 is stereo (most speakers are stereo).
	op.ChannelCount = 2

	// Format of the source. go-mp3's format is signed 16bit integers.
	op.Format = oto.FormatSignedInt16LE

	// Remember that you should **not** create more than one context
	//var otoCtx *oto.Context
	var readyChan chan struct{}
	otoCtx, readyChan, err = oto.NewContext(op)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	// It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
	<-readyChan

	// Create a new 'player' that will handle our sound. Paused by default.
	//var player *oto.Player
	systemNotificationPlayer = otoCtx.NewPlayer(decodedMp3)

}

func initiateInvalidNotification() {

	var err error

	// Convert the pure bytes into a reader object that can be used with the mp3 decoder
	var fileBytesReader *bytes.Reader
	fileBytesReader = bytes.NewReader(invalidSelectionAsByteArray)

	// Decode file
	var decodedMp3 *mp3.Decoder
	decodedMp3, err = mp3.NewDecoder(fileBytesReader)
	if err != nil {
		panic("mp3.NewDecoder failed: " + err.Error())
	}

	// Prepare an Oto context (this will use your default audio device) that will
	// play all our sounds. Its configuration can't be changed later.

	var op *oto.NewContextOptions
	op = &oto.NewContextOptions{}

	// Usually 44100 or 48000. Other values might cause distortions in Oto
	op.SampleRate = 44100

	// Number of channels (aka locations) to play sounds from. Either 1 or 2.
	// 1 is mono sound, and 2 is stereo (most speakers are stereo).
	op.ChannelCount = 2

	// Format of the source. go-mp3's format is signed 16bit integers.
	op.Format = oto.FormatSignedInt16LE

	// Remember that you should **not** create more than one context
	//var otoCtx *oto.Context
	var readyChan chan struct{}
	otoCtx, readyChan, err = oto.NewContext(op)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	// It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
	<-readyChan

	// Create a new 'player' that will handle our sound. Paused by default.
	//var player *oto.Player
	invalidNotificationPlayer = otoCtx.NewPlayer(decodedMp3)

}

// Play the System Notification Sound
func playSystemNotification() {

	var err error

	// Initiate sound engine if that not has been done
	initSoundEngine()

	// Restart from the beginning (or go to any location in the sound) using seek
	//_, err = systemNotificationPlayer.(io.Seeker).Seek(0, io.SeekStart) //newPos
	_, err = systemNotificationPlayer.Seek(0, io.SeekStart) //newPos
	if err != nil {
		panic("player.Seek failed: " + err.Error())
	}

	// Play starts playing the sound and returns without waiting for it (Play() is async).
	systemNotificationPlayer.Play()

	// We can wait for the sound to finish playing using something like this
	//for systemNotificationPlayer.IsPlaying() {
	//	time.Sleep(time.Millisecond)
	//}

}

// Play the Invalid Notification Sound
func playInvalidNotification() {

	var err error

	// Initiate sound engine if that not has been done
	initSoundEngine()

	// Restart from the beginning (or go to any location in the sound) using seek
	//_, err = systemNotificationPlayer.(io.Seeker).Seek(0, io.SeekStart) //newPos
	_, err = invalidNotificationPlayer.Seek(0, io.SeekStart) //newPos
	if err != nil {
		panic("player.Seek failed: " + err.Error())
	}

	// Play starts playing the sound and returns without waiting for it (Play() is async).
	invalidNotificationPlayer.Play()

	// We can wait for the sound to finish playing using something like this
	//for systemNotificationPlayer.IsPlaying() {
	//	time.Sleep(time.Millisecond)
	//}

}
