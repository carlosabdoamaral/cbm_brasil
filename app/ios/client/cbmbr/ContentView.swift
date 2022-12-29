//
//  ContentView.swift
//  cbmbr
//
//  Created by Carlos Amaral on 25/12/22.
//

import SwiftUI

struct ContentView: View {
    var body: some View {
        TabView {
            HomeView()
                .tabItem {
                    Image(systemName: "house")
                    Text("Menu")
                }
            
            MapView()
                .tabItem {
                    Image(systemName: "map.fill")
                    Text("Mapa")
                }
            
            HowToView()
                .tabItem {
                    Image(systemName: "cross.case.fill")
                    Text("Conhecimento")
                }
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
