//
//  HowToDetailsView.swift
//  cbmbr
//
//  Created by Carlos Amaral on 25/12/22.
//

import SwiftUI
import AVKit

struct HowToDetailsView: View {
    @State var item : HowToItemModel
    @State private var currentStepIndex : Int = 0
    @State private var currentStep : HowToStepModel = HowToTemplateData[0].steps[0]
    @State private var isShowingAlert : Bool = false
    @State private var player = AVPlayer()
    
    let defaultHorizontalPadding : CGFloat = 20.0
    
    func handleNextStep() {
        item.steps[currentStepIndex].didCheck = true
        
        currentStepIndex+=1
        currentStep = item.steps[currentStepIndex]
    }
    
    /// TODO: Chamar o método
    func handleDismissAlert() {
        isShowingAlert = false
    }
    
    func completeCourse() {
        item.steps[currentStepIndex].didCheck = true
        isShowingAlert = true
    }
    
    func isLastStep() -> Bool {
        return currentStepIndex == item.steps.count - 1
    }
    
    var body: some View {
        ZStack {
            Color.secondary.edgesIgnoringSafeArea(.all).opacity(0.1)
            
            VStack(alignment: .leading) {
                VideoPlayer(player: AVPlayer(
                    url: Bundle.main.url(
                        forResource: "template_video",
                        withExtension: "mp4"
                    )!
                ))
                .frame(height: 250)
                
                HStack {
                    VStack(alignment: .leading) {
                        Text("Atual:")
                            .font(.subheadline)
                            .foregroundColor(.secondary)
                        Text(currentStep.title)
                            .font(.title)
                            .bold()
                    }
                    
                    Spacer()
                    
                    if !isLastStep() {
                        Button {
                            handleNextStep()
                        } label: {
                            Image(systemName: "arrow.forward.circle.fill")
                                .frame(width: 30, height: 30)
                        }
                    } else {
                        Button {
                            completeCourse()
                        } label: {
                            Image(systemName: "trophy.fill")
                                .frame(width: 30, height: 30)
                        }
                    }
                }
                .padding(.top, 16)
                .padding(.horizontal, self.defaultHorizontalPadding)
                .alert(isPresented: $isShowingAlert) {
                    Alert(
                        title: Text("Status"),
                        message: Text("Parabéns! Você concluiu o tutorial: \(item.title)"),
                        dismissButton:.default(Text("Fechar"))
                    )
                }
                
                ScrollView {
                    ForEach(item.steps) { step in
                        HStack {
                            Text(step.title)
                            Spacer()
                            
                            if step.didCheck {
                                Image(systemName: "checkmark.circle.fill")
                                    .foregroundColor(.blue)
                            }
                        }
                        .frame(height: 40)
                        
                        Divider()
                    }
                    .padding(.horizontal, self.defaultHorizontalPadding)
                }
                
                Spacer()
            }
        }
        .navigationTitle($item.title)
        .navigationBarTitleDisplayMode(.inline)
    }
}

struct HowToDetailsView_Previews: PreviewProvider {
    static var previews: some View {
        HowToDetailsView(item: HowToTemplateData[0])
    }
}
