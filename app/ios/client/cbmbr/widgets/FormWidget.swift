//
//  FormWidget.swift
//  cbmbr
//
//  Created by Carlos Amaral on 25/12/22.
//

import SwiftUI

struct LargeSecondaryButton : View {
    @State private var isFocused : Bool = false
    let title : String
    let titleOnFocus : String
    let action : () -> Void
    
    var body : some View {
        HStack {
            if isFocused {
                Text(titleOnFocus)
            } else {
                Text(title)
            }
            
            Spacer()
        }
        .onTapGesture {
            if !self.titleOnFocus.isEmpty {
                isFocused.toggle()
            }
            
            action()
        }
        .foregroundColor(.accentColor)
        .frame(height: 60)
        .padding(.leading)
        .background(.primary.opacity(0.08))
        .cornerRadius(10)
    }
}

struct LargePrimaryButton : View {
    let title : String
    let action : () -> Void
    
    var body : some View {
        HStack {
            Spacer()
            Text(title)
            Spacer()
        }
        .onTapGesture { action() }
        .foregroundColor(.white)
        .frame(height: 60)
        .padding(.leading)
        .background(Color.accentColor)
        .cornerRadius(10)
        .bold()
        .font(.title3)
    }
}

struct CustomTextField : View {
    let placeholder : String
    @State var title : Binding<String>
    
    var body: some View {
        ZStack {
            TextField(self.placeholder, text: title)
                .padding()
        }
        .frame(height: 60)
        .background(.primary.opacity(0.08))
        .cornerRadius(10)
    }
}


struct FormWidget_Previews: PreviewProvider {
    static var previews: some View {
        ScrollView {
            LargeSecondaryButton(title: "Exemplo", titleOnFocus: "Exemplo focus") {  }
            LargePrimaryButton(title: "Exemplo") {  }
        }
        .padding(.all)
    }
}
