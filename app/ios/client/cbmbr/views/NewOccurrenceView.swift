//
//  NewOccurrenceView.swift
//  cbmbr
//
//  Created by Carlos Amaral on 25/12/22.
//

import SwiftUI
import MapKit
import AVKit
import CoreLocation
import CoreLocationUI

struct NewOccurrenceModel : Identifiable {
    let id : UUID = UUID()
    let title : String = ""
    let description : String = ""
    let image : UIImage? = UIImage(systemName: "hammer")
    let location : Location = mapMarkers[0]
}


// MARK: Erros:
// - Salvar o audio no diretorio de documentos
struct NewOccurrenceView: View {
    @State var title : String = ""
    
    @State var isRecordingAudio : Bool = false
    @State var audioRecorderSession : AVAudioSession!
    @State var audioRecorder : AVAudioRecorder!
    @State var isShowingAudioPermissionAlert : Bool = false
    @State var audios : [URL] = []
    
    @State var cameraIsOpen : Bool = false
    @State var image : Image? = nil
    let imageSize : CGFloat = 360.0
    
    @ObservedObject var locationManager = LocationManager()
    @State var locationInputValue = ""
    
    func handleRequestAudio() {
        do {
            self.audioRecorderSession = AVAudioSession.sharedInstance()
            try self.audioRecorderSession.setCategory(.playAndRecord)
            
            self.audioRecorderSession.requestRecordPermission { didAccept in
                if didAccept {
                    getAudios()
                } else {
                    self.isShowingAudioPermissionAlert = true
                }
            }
            
        } catch {
            print(error.localizedDescription)
        }
    }
    
    var body: some View {
        ZStack {
            ScrollView {
                VStack(alignment: .leading) {
                    Text("Nos conte o que está acontecendo e iremos enviar um suporte para você")
                        .font(.subheadline)
                        .foregroundColor(.secondary)
                        .padding(.bottom, 30)
                    
                    image?
                        .resizable()
                        .scaledToFit()
                        .frame(width: self.imageSize, height: self.imageSize)
                        .cornerRadius(10)
                    
                    CustomTextField(placeholder: "Insira um título", title: $title)
                    CustomTextField(placeholder: "Sua localização", title: $locationInputValue)
                    
                    LargeSecondaryButton(title: "Gravar áudio", titleOnFocus: "Gravando...") {
                        isRecordingAudio.toggle()
                        handleRecordAudio()
                    }
                    
                    LargeSecondaryButton(title: "Tirar foto", titleOnFocus: "") {
                        cameraIsOpen.toggle()
                    }
                    
                    LargePrimaryButton(title: "Enviar") {
                        
                    }
                    Spacer()
                }
                .padding(.horizontal)
            }
            
            if cameraIsOpen {
                CaptureImageView(isShown: $cameraIsOpen, image: $image)
                    .edgesIgnoringSafeArea(.all)
            }
        }
        .navigationTitle("Criar")
        .alert(isPresented: $isShowingAudioPermissionAlert) {
            Alert(
                title: Text("Status"),
                message: Text("Você deve permitir o uso do microfone para usar essa funcionalidade"),
                dismissButton: .default(Text("Fechar"))
            )
        }
        //        .onAppear {
        //            self.handleRequestAudio()
        //            let coordinates = getUserCoordinates(locationManager: locationManager)
        //
        //            switch locationManager.statusString {
        //            case "authorizedWhenInUse":
        //                locationInputValue = "\(coordinates!.latitude) \(coordinates!.longitude)"
        //                break
        //            case "authorizedAlways":
        //                locationInputValue = "\(coordinates!.latitude) \(coordinates!.longitude)"
        //                break
        //            default:
        //                locationInputValue = "Insira a localização"
        //            }
        //        }
    }
}

struct NewOccurrenceView_Previews: PreviewProvider {
    static var previews: some View {
        NewOccurrenceView()
    }
}
