//------------------------------------------------------------------------------
// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: osa.fbe
// FBE version: 1.10.0.0
//------------------------------------------------------------------------------

import Foundation
import Fbe

public protocol ExtraBase {
    var name: String { get set }
    var detail: String { get set }
    var sex: Sex { get set }
    var flag: MyFLags { get set }
}

public protocol ExtraInheritance {
    var parent: Extra { get set }
}

extension ExtraInheritance {
    public var name: String {
        get { return parent.name }
        set { parent.name = newValue }
    }
    public var detail: String {
        get { return parent.detail }
        set { parent.detail = newValue }
    }
    public var sex: Sex {
        get { return parent.sex }
        set { parent.sex = newValue }
    }
    public var flag: MyFLags {
        get { return parent.flag }
        set { parent.flag = newValue }
    }
}

public struct Extra: ExtraBase, Comparable, Hashable, Codable {
    public var name: String = ""
    public var detail: String = ""
    public var sex: Sex = Osa.Sex()
    public var flag: MyFLags = Osa.MyFLags()

    public init() { }
    public init(name: String, detail: String, sex: Sex, flag: MyFLags) {

        self.name = name
        self.detail = detail
        self.sex = sex
        self.flag = flag
    }

    public init(other: Extra) {
        self.name = other.name
        self.detail = other.detail
        self.sex = other.sex
        self.flag = other.flag
    }

    public init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        name = try container.decode(String.self, forKey: .name)
        detail = try container.decode(String.self, forKey: .detail)
        sex = try container.decode(Osa.Sex.self, forKey: .sex)
        flag = try container.decode(Osa.MyFLags.self, forKey: .flag)
    }

    public func clone() throws -> Extra {
        // Serialize the struct to the FBE stream
        let writer = ExtraModel()
        try _ = writer.serialize(value: self)

        // Deserialize the struct from the FBE stream
        let reader = ExtraModel()
        reader.attach(buffer: writer.buffer)
        return reader.deserialize()
    }

    public static func < (lhs: Extra, rhs: Extra) -> Bool {
        return true
    }

    public static func == (lhs: Extra, rhs: Extra) -> Bool {
        return true
    }

    public func hash(into hasher: inout Hasher) {
    }

    public var description: String {
        var sb = String()
        sb.append("Extra(")
        sb.append("name="); sb.append("\""); sb.append(name); sb.append("\"")
        sb.append(",detail="); sb.append("\""); sb.append(detail); sb.append("\"")
        sb.append(",sex="); sb.append(sex.description)
        sb.append(",flag="); sb.append(flag.description)
        sb.append(")")
        return sb
    }
    private enum CodingKeys: String, CodingKey {
        case name
        case detail
        case sex
        case flag
    }

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encode(name, forKey: .name)
        try container.encode(detail, forKey: .detail)
        try container.encode(sex, forKey: .sex)
        try container.encode(flag, forKey: .flag)
    }

    public func toJson() throws -> String {
        return String(data: try JSONEncoder().encode(self), encoding: .utf8)!
    }

    public static func fromJson(_ json: String) throws -> Extra {
        return try JSONDecoder().decode(Extra.self, from: json.data(using: .utf8)!)
    }
}
