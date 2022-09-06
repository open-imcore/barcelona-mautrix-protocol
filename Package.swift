// swift-tools-version: 5.4

import PackageDescription

let package = Package(
    name: "BarcelonaMautrixIPCProtobuf",
    products: [
        .library(name: "BarcelonaMautrixIPCProtobuf", targets: ["BarcelonaMautrixIPCProtobuf"])
    ],
    dependencies: [
        .package(name: "SwiftProtobuf", url: "https://github.com/apple/swift-protobuf", from: "1.20.0"),
        .package(name: "GRPC", url: "https://github.com/grpc/grpc-swift.git", from: "1.9.0")
    ],
    targets: [
        .target(name: "BarcelonaMautrixIPCProtobuf", dependencies: [
            "SwiftProtobuf", "GRPC"
        ])
    ]
)