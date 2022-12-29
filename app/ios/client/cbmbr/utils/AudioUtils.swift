//
//  AudioUtils.swift
//  cbmbr
//
//  Created by Carlos Amaral on 25/12/22.
//

import Foundation
import AVKit

func getAudios() {
    print("[*] Get Audios")
    
    do {
        let url = FileManager.default.urls(for: .documentDirectory, in: .userDomainMask)[0]
        let result = try FileManager.default.contentsOfDirectory(at: url, includingPropertiesForKeys: nil, options: .producesRelativePathURLs)

        print(result.count)
        
    } catch {
        print(error.localizedDescription)
    }
}

func handleRecordAudio() {
    var audioRecorder : AVAudioRecorder? = nil
    
    do {
        // MARK: Handle Record Audio
        // We are going to store audio in document directory
        // it will always save with the same name, this way
        // will update the last audio with the latest record
        
        _ = getDocumentsDirectory().appendingPathComponent("occurrenceAudio.m4a")
        
        // or if you want to have multiples audios
        // _ = url.appendingPathComponent("occurrenceAudio#\(self.audios.count + 1).m4a")
        
        
        let settings = [
            AVFormatIDKey :  Int(kAudioFormatMPEG4AAC),
            AVSampleRateKey : 12000,
            AVNumberOfChannelsKey : 1,
            AVEncoderAudioQualityKey: AVAudioQuality.high.rawValue
        ]
        
        audioRecorder = try AVAudioRecorder(url: getDocumentsDirectory(), settings: settings)
        audioRecorder!.record()
    } catch {
        print(error.localizedDescription)
    }
}

func handleStopRecordAudio(audioRecorder : AVAudioRecorder) {
    audioRecorder.stop()
    getAudios()
    return
}
