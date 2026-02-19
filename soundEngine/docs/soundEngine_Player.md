# soundEngine_Player.go

## File Overview
- Path: `soundEngine/soundEngine_Player.go`
- Package: `soundEngine`
- Functions/Methods: `10`
- Imports: `8`

## File Purpose
No concise file-level comment detected. Purpose inferred from declarations below.

## Potential Entry Points
- `CloseDownSoundEngine`
- `InitSoundEngine`

## Imports
- `FenixTesterGui/common_code`
- `bytes`
- `embed`
- `github.com/ebitengine/oto/v3`
- `github.com/hajimehoshi/go-mp3`
- `github.com/sirupsen/logrus`
- `io`
- `time`

## Declared Types
- `SoundType`

## Declared Constants
- `InvalidNotificationSound`
- `SystemNotificationSound`
- `UserNeedToRespondSound`
- `soundChannelSize`

## Declared Variables
- `PlaySoundChannel`
- `invalidNotificationPlayer`
- `invalidSelectionAsByteArray`
- `otoCtx`
- `systemNotificationAsByteArray`
- `systemNotificationPlayer`
- `userNeedToRespondAsByteArray`
- `userNeedToRespondPlayer`

## Functions and Methods
### playerChannelReader
- Signature: `func playerChannelReader()`
- Exported: `false`
- Control-flow features: `for/range, switch`
- Doc: The reader for the Sound player
- Internal calls: `playSystemNotification`, `playInvalidNotification`, `playUserNeedToRespond`

### initiatePlayerChannelEngine
- Signature: `func initiatePlayerChannelEngine()`
- Exported: `false`
- Control-flow features: `go`
- Doc: Initiate PlayerChannelEngine
- Internal calls: `playerChannelReader`

### InitSoundEngine
- Signature: `func InitSoundEngine()`
- Exported: `true`
- Control-flow features: `if`
- Doc: Init the Sound Enigine if that hasn't been done
- Internal calls: `initiatePlayerChannelEngine`, `initiateSystemNotification`, `initiateInvalidNotification`, `initiateUserNeedToRespond`
- Selector calls: `oto.NewContext`, `err.Error`

### CloseDownSoundEngine
- Signature: `func CloseDownSoundEngine()`
- Exported: `true`
- Control-flow features: `if`
- Selector calls: `systemNotificationPlayer.Close`, `err.Error`, `invalidNotificationPlayer.Close`

### initiateSystemNotification
- Signature: `func initiateSystemNotification()`
- Exported: `false`
- Control-flow features: `if`
- Selector calls: `bytes.NewReader`, `mp3.NewDecoder`, `err.Error`, `otoCtx.NewPlayer`

### initiateInvalidNotification
- Signature: `func initiateInvalidNotification()`
- Exported: `false`
- Control-flow features: `if`
- Selector calls: `bytes.NewReader`, `mp3.NewDecoder`, `err.Error`, `otoCtx.NewPlayer`

### initiateUserNeedToRespond
- Signature: `func initiateUserNeedToRespond()`
- Exported: `false`
- Control-flow features: `if`
- Selector calls: `bytes.NewReader`, `mp3.NewDecoder`, `err.Error`, `otoCtx.NewPlayer`

### playSystemNotification
- Signature: `func playSystemNotification()`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Play the System Notification Sound
- Selector calls: `systemNotificationPlayer.Seek`, `err.Error`, `systemNotificationPlayer.Play`, `systemNotificationPlayer.IsPlaying`, `time.Sleep`

### playInvalidNotification
- Signature: `func playInvalidNotification()`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Play the Invalid Notification Sound
- Selector calls: `invalidNotificationPlayer.Seek`, `err.Error`, `invalidNotificationPlayer.Play`, `systemNotificationPlayer.IsPlaying`, `time.Sleep`

### playUserNeedToRespond
- Signature: `func playUserNeedToRespond()`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Play the User Need to Respond Sound
- Selector calls: `invalidNotificationPlayer.Seek`, `err.Error`, `userNeedToRespondPlayer.Play`, `systemNotificationPlayer.IsPlaying`, `time.Sleep`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
