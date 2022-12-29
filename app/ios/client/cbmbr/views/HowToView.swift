//
//  HowToView.swift
//  cbmbr
//
//  Created by Carlos Amaral on 25/12/22.
//

import SwiftUI

struct HowToView: View {
    @State private var searchInputValue : String = ""
    @State private var isShowingPageInfos : Bool = false
    
    var searchResults : [HowToItemModel] {
        if searchInputValue.isEmpty {
            return HowToTemplateData
        } else {
            var res : [HowToItemModel] = []
            
            HowToTemplateData.forEach { row in
                if  row.title.uppercased().contains(searchInputValue.uppercased()) ||
                        row.description.uppercased().contains(searchInputValue.uppercased()) {
                    res.append(row)
                }
            }
            
            return res
        }
    }

    let rowHeight : CGFloat = 80.0
    
    var body: some View {
        NavigationView {
            List {
                ForEach(searchResults) { item in
                    NavigationLink {
                        HowToDetailsView(item: item)
                    } label: {
                        HStack {
                            VStack(alignment: .leading) {
                                Text(item.title)
                                    .font(.headline)
                                    .foregroundColor(.primary)
                                
                                Text(item.description)
                                    .font(.subheadline)
                                    .multilineTextAlignment(.leading)
                                
                                Spacer()
                            }
                            .padding(.top, 15)
                            .foregroundColor(.secondary)
                            
                            Spacer()
                        }
                    }
                }
            }
            .searchable(text: $searchInputValue, prompt: "Buscar tutorial")
            .navigationTitle("Conhecimento")
            .alert(isPresented: $isShowingPageInfos) {
                Alert(
                    title: Text("Sobre"),
                    message: Text("Na aba de conhecimento, você pode aprender algumas técnicas de pronto socorro que podem auxiliar em momento de emergência"),
                    dismissButton: .default(Text("Ok"))
                )
            }
            .toolbar {
                Button {
                    self.isShowingPageInfos.toggle()
                } label: {
                    Image(systemName: "questionmark.circle")
                }
            }
        }
    }
}

struct HowToView_Previews: PreviewProvider {
    static var previews: some View {
        HowToView()
    }
}
