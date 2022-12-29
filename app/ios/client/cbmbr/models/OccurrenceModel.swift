//
//  OccurrenceModel.swift
//  cbmbr
//
//  Created by Carlos Amaral on 25/12/22.
//

import Foundation
import SwiftUI
import MapKit

struct OccurrenceModel : Identifiable {
    let id : UUID = UUID()
    var title : String
    var location : CLLocationCoordinate2D
    var image : Image
    
}
