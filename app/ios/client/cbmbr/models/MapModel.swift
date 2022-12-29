//
//  MapModel.swift
//  cbmbr
//
//  Created by Carlos Amaral on 25/12/22.
//

import Foundation
import MapKit

struct Location : Identifiable {
    let id = UUID()
    let name : String
    let coordinate : CLLocationCoordinate2D
}
