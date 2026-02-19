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
### CloseDownSoundEngine
- Signature: `func CloseDownSoundEngine()`
- Exported: `true`
- Control-flow features: `if`
- Selector calls: `err.Error`, `invalidNotificationPlayer.Close`, `systemNotificationPlayer.Close`

### InitSoundEngine
- Signature: `func InitSoundEngine()`
- Exported: `true`
- Control-flow features: `if`
- Doc: Init the Sound Enigine if that hasn't been done
- Internal calls: `initiateInvalidNotification`, `initiatePlayerChannelEngine`, `initiateSystemNotification`, `initiateUserNeedToRespond`
- Selector calls: `err.Error`, `oto.NewContext`

### initiateInvalidNotification
- Signature: `func initiateInvalidNotification()`
- Exported: `false`
- Control-flow features: `if`
- Selector calls: `bytes.NewReader`, `err.Error`, `mp3.NewDecoder`, `otoCtx.NewPlayer`

### initiatePlayerChannelEngine
- Signature: `func initiatePlayerChannelEngine()`
- Exported: `false`
- Control-flow features: `go`
- Doc: Initiate PlayerChannelEngine
- Internal calls: `playerChannelReader`

### initiateSystemNotification
- Signature: `func initiateSystemNotification()`
- Exported: `false`
- Control-flow features: `if`
- Selector calls: `bytes.NewReader`, `err.Error`, `mp3.NewDecoder`, `otoCtx.NewPlayer`

### initiateUserNeedToRespond
- Signature: `func initiateUserNeedToRespond()`
- Exported: `false`
- Control-flow features: `if`
- Selector calls: `bytes.NewReader`, `err.Error`, `mp3.NewDecoder`, `otoCtx.NewPlayer`

### playInvalidNotification
- Signature: `func playInvalidNotification()`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Play the Invalid Notification Sound
- Selector calls: `err.Error`, `invalidNotificationPlayer.Play`, `invalidNotificationPlayer.Seek`, `systemNotificationPlayer.IsPlaying`, `time.Sleep`

### playSystemNotification
- Signature: `func playSystemNotification()`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Play the System Notification Sound
- Selector calls: `err.Error`, `systemNotificationPlayer.IsPlaying`, `systemNotificationPlayer.Play`, `systemNotificationPlayer.Seek`, `time.Sleep`

### playUserNeedToRespond
- Signature: `func playUserNeedToRespond()`
- Exported: `false`
- Control-flow features: `if, for/range`
- Doc: Play the User Need to Respond Sound
- Selector calls: `err.Error`, `invalidNotificationPlayer.Seek`, `systemNotificationPlayer.IsPlaying`, `time.Sleep`, `userNeedToRespondPlayer.Play`

### playerChannelReader
- Signature: `func playerChannelReader()`
- Exported: `false`
- Control-flow features: `for/range, switch`
- Doc: The reader for the Sound player
- Internal calls: `playInvalidNotification`, `playSystemNotification`, `playUserNeedToRespond`

## Behavioral Summary
This file summary is generated from AST analysis. For exact runtime behavior (ordering, side effects, retries, failure semantics), validate against source and tests.
