//
//  SystemUtils.swift
//  cbmbr
//
//  Created by Carlos Amaral on 25/12/22.
//

import Foundation

func getDocumentsDirectory() -> URL {
    return FileManager.default.urls(for: .documentDirectory, in: .userDomainMask)[0]
}
