FROM julesreyn/flutter-android-sdk:jdk17-android33-flutter3.24.3

WORKDIR /home/builder/app

USER root

COPY . .

RUN chown -R builder:builder /home/builder/app

COPY pubspec.yaml .

USER builder

RUN flutter clean && flutter pub get
RUN flutter packages get
RUN flutter build apk --release

USER root

CMD ["sh", "-c",  "cp build/app/outputs/flutter-apk/app-release.apk /home/builder/app/mobile/client.apk"]
