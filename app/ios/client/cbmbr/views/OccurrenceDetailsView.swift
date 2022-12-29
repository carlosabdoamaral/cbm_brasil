//
//  OccurrenceDetailsView.swift
//  cbmbr
//
//  Created by Carlos Amaral on 28/12/22.
//

import SwiftUI

struct OccurrenceDetailsView: View {
    let infos : OccurrenceModel
    var body: some View {
        ScrollView {
            VStack(alignment: .leading) {
                VStack(alignment: .leading) {
                    Text("Título")
                        .font(.subheadline)
                        .foregroundColor(.secondary)
                    Text(infos.title)
                        .font(.title)
                        .bold()
                        .lineLimit(1)
                    
                    infos.image
                        .resizable()
                        .frame(height: 300)
                        .padding()
                        .background(.primary.opacity(0.08))
                        .cornerRadius(10)
                    
                    LargeSecondaryButton(title: "Ouvir áudio", titleOnFocus: "Parar de ouvir") {
                        // code
                    }
                    
                    LargeSecondaryButton(title: "\(infos.location.longitude), \(infos.location.latitude)", titleOnFocus: "\(infos.location.longitude), \(infos.location.latitude)") {
                        // code
                    }
                }
                
                Spacer()
            }
//            .navigationTitle("")
            .padding()
            .navigationBarTitleDisplayMode(.inline)
            .toolbar {
                Button {
                    
                } label: {
                    Image(systemName: "exclamationmark.triangle")
                }
                .foregroundColor(.red)
            }
        }
    }
}

struct OccurrenceDetailsView_Previews: PreviewProvider {
    static var previews: some View {
        OccurrenceDetailsView(infos: OccurrencesTemplateData[0])
    }
}
