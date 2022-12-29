//
//  HowToModel.swift
//  cbmbr
//
//  Created by Carlos Amaral on 25/12/22.
//

import Foundation

struct HowToItemModel : Identifiable {
    let id : UUID = UUID()
    var title : String
    var description : String
    var steps : [HowToStepModel]
}

struct HowToStepModel : Identifiable {
    var id : UUID = UUID()
    var title : String
    var description : String
    var didCheck : Bool
}
