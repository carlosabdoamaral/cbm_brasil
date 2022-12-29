//
//  TemplateData.swift
//  cbmbr
//
//  Created by Carlos Amaral on 25/12/22.
//

import Foundation
import SwiftUI
import MapKit

let HowToTemplateData : [HowToItemModel] = [
    HowToItemModel(title: "Teste 01", description: "lalall lal lall lallalal alallala llalala", steps: [
        HowToStepModel(title: "Passo 01", description: "Abrir os olhos", didCheck: false),
        HowToStepModel(title: "Passo 02", description: "Fechar os olhos", didCheck: false),
        HowToStepModel(title: "Passo 03", description: "Acordar", didCheck: false)
    ]),
    
    HowToItemModel(title: "Respiraçao", description: "lalall lal lall lallalal alallala llalala", steps: [
        HowToStepModel(title: "Passo 01", description: "Abrir os olhos", didCheck: false),
        HowToStepModel(title: "Passo 02", description: "Fechar os olhos", didCheck: false),
        HowToStepModel(title: "Passo 03", description: "Acordar", didCheck: false)
    ])
]

let OccurrencesTemplateData : [OccurrenceModel] = [
    OccurrenceModel(title: "Teste", location: CLLocationCoordinate2D(latitude: 51.508, longitude: -0.076), image: Image(systemName: "hammer"))
]
