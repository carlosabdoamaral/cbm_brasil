//
//  MapView.swift
//  cbmbr
//
//  Created by Carlos Amaral on 25/12/22.
//

import SwiftUI

struct MapView: View {
    var body: some View {
        NavigationView {
            ZStack {
                MapWidget()
                    .edgesIgnoringSafeArea([.top, .horizontal])
            }
            .navigationTitle("Mapa")
        }
    }
}

struct MapView_Previews: PreviewProvider {
    static var previews: some View {
        MapView()
    }
}
