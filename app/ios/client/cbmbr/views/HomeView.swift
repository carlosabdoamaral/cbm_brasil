//
//  HomeView.swift
//  cbmbr
//
//  Created by Carlos Amaral on 25/12/22.
//

import SwiftUI

struct HomeView: View {
    var occurrences : [OccurrenceModel] = OccurrencesTemplateData
    
    var body: some View {
        NavigationView {
            ZStack {
                ScrollView {
                    VStack(alignment: .leading) {
                        SectionTitleWidget(title: "Ocorrências", secondaryButtonText: "Criar", secondaryButtonNextView: NewOccurrenceView())
                        
                        ForEach(occurrences) { row in
                            NavigationLink {
                                OccurrenceDetailsView(infos: row)
                            } label: {
                                HStack {
                                    VStack(alignment: .leading) {
                                        Text(row.title)
                                            .font(.headline)
                                        Text("\(row.location.longitude) - \(row.location.latitude)")
                                            .foregroundColor(.secondary)
                                            .font(.subheadline)
                                    }
                                    Spacer()
                                }
                            }
                            .padding(.all)
                            .background(Color.primary.opacity(0.05))
                            .cornerRadius(10)
                        }
                        .listStyle(.plain)
                        
                        Spacer()
                    }
                }
            }
            .padding(.horizontal)
            .navigationTitle("Menu")
            .toolbar {
                NavigationLink(destination: NewOccurrenceView()) {
                    Image(systemName: "plus.circle")
                        .foregroundColor(.red)
                        .font(.title2)
                        .bold()
                }
            }
        }
    }
}

struct SectionTitleWidget: View {
    let title : String
    let secondaryButtonText : String
    let secondaryButtonNextView : any View
    
    var body : some View {
        HStack {
            Text(self.title)
                .font(.title2)
                .bold()
            
            Spacer()
            
            NavigationLink(destination: NewOccurrenceView()) {
                Text(self.secondaryButtonText)
            }
        }
        .padding(.vertical)
    }
}

struct HomeView_Previews: PreviewProvider {
    static var previews: some View {
        HomeView()
    }
}
