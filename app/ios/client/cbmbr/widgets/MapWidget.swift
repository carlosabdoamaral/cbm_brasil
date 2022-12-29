//
//  MapWidget.swift
//  cbmbr
//
//  Created by Carlos Amaral on 25/12/22.
//

import SwiftUI
import MapKit

struct MapWidget: View {
    @ObservedObject var locationManager = LocationManager()
    
    @State private var mapRegion = MKCoordinateRegion(
        center: CLLocationCoordinate2D(latitude: 51.5, longitude: -0.12),
        span: MKCoordinateSpan(latitudeDelta: 0.005, longitudeDelta: 0.005)
    )
    
    var body: some View {
        Map(coordinateRegion: $mapRegion, annotationItems: mapMarkers) { location in
            MapAnnotation(coordinate: location.coordinate) {
                NavigationLink {
                    Text(location.name)
                } label: {
                    Circle()
                        .stroke(.red, lineWidth: 3)
                        .frame(width: 44, height: 44)
                }
            }
        }
        .onAppear {
            let coordinates = getUserCoordinates(locationManager: locationManager)
            
            switch locationManager.statusString {
            case "authorizedWhenInUse":
                mapRegion.center = coordinates!
                break
            case "authorizedAlways":
                mapRegion.center = coordinates!
                break
            default:
                break
            }
        }
    }
}

struct MapWidget_Previews: PreviewProvider {
    static var previews: some View {
        MapWidget()
    }
}
