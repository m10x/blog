---
title: "HackTheBox Invite Code WriteUp | Tipps + Anleitung"
date: 2021-03-09T09:21:56+01:00
toc: false
images:
tags:
  - hackthebox
---

Tipps und Anleitung dazu, wie man an den Registrierungs Code für HackTheBox kommt.

## Einleitung

HackTheBox ist eine online Plattform bei der man seine IT-Sicherheit und Penetration Test Fähigkeiten anwenden, testen und verbessern kann.
Um sich registrieren zu können muss man einen kleinen Test bestehen. Ich habe diesen Eintrag geschrieben, da der Test kein Hindernis sein soll, sich bei HackTheBox zu beteiligen.
Zum Beispiel, wenn man nur "hineinschnuppern" möchte oder bei einer Sache nicht weiter kommt.
Allerdings gilt wie immer beim PenTesting das Motto "Try Harder!". :)

## Tipps

- Untersuche den Quellcode.
- Du wirst die Console vom Browser benötigen um einen Java Script Befehl auszuführen
- Du musst Base64 und eventuell ROT13 dekodieren.
- Um einen POST-request zu senden kannst du Browser Addons benutzen

## Video

[Kurzes Video Walkthrough ohne Erklärungen](https://www.youtube.com/watch?v=5bEOmzgnWC4)

## Anleitung
Als erstes untersuchen wir den Quellcode (F12 oder Rechtsklick und (Element) untersuchen)  
![Element untersuchen][picture1]  
Können wir dort etwas interessantes finden?  
![inviteapi.min.js][picture2]  
Das JavaScript mit dem Pfad /js/inviteapi.min.js sieht vielversprechend aus oder?  
Sehen wir uns es mal genauer an...  
![makeInviteCode][picture3]  
POST und makeInviteCode sehen nützlich aus. POST bedeutet in diesem Fall, dass das JavaScript ein HTTP Post Request unterstützt. makeInviteCode ist eine Funktion des JavaScripts.  
Was können wir nun mit diesem Wissen anfagen?  
![makeInviteCode()][picture4]  
Wir führen die die Funktion makeInviteCode() in der Konsole unseres Internetbrowsers aus (in meinem Fall Firefox Quantum Developer Edition) und erhalten folgendes:  
data = Va beqre gb trarengr gur vaivgr pbqr, znxr n CBFG erdhrfg gb /ncv/vaivgr/trarengr  
enctype = ROT13  
Wir haben nun also einen String der mit Hilfe von "ROT13" verschlüsselt wurde. "ROT13" ist eine Caesar-Verschlüsselung, bei der alle Buchstaben um 13 Stellen im Alphabet verschoben werden. Wer wissen möchte, wie man mit "ROT13" verschlüsselte Strings ganz einfach mit Bash oder Python entschlüsseln kann, kann dies in meinem "OverTheWire Bandit" Anleitung nachsehen, bei Level 11-12.
Der Einfachheit halber kann man auch eine Internetseite benutzen, die das auf Knopfdruck für einen macht oder ein Browser-Plugin. Ich benutze das Mozilla Firefox Quantum Plugin "New Hackbar" dafür.

Der entschlüsselte Text ist wie folgt: In order to generate the invite code, make a POST request to /api/invite/generate

Um sehr einfach ein POST request zu machen, kann man ein Browser-Plugin dafür benutzen.
Firefox Quantum: New Hackbar
Firefox Älter: Hackbar
Google Chrome: Postman
Ich werde Firefox Quantum mit dem "New Hackbar" Plugin dafür benutzen, was du benutzt, ist natürlich dir überlassen.

Wir sehen nun code = SIJFVEwtSktSRkktSFIPUEYtTFIGR08tTU5QVVM= und format = encoded
Wir haben jetzt also anscheinend, den Code den wir für die Registrierung benötigen, aber dieser ist noch codiert ( = encoded).
Allerdings steht dort nicht, wie der Code codiert wurde. Wie können wir dies herausfinden? Trial & Error?

Der String code hat an letzter Stelle ein =. Deshalb ist die Wahrscheinlichkeit groß, dass der String mit Hilfe von Base64 codiert wurde,
da eine Base64 Kodierung immer mit einem = endet.
Versuchen wir mal unser Glück.

Geschafft!! Wir haben nun den Code, welchen wir benötigt haben.

Probieren wir ihn mal aus. :)

Es hat funktioniert!

[picture1]: https://i.imgur.com/H11xz1Q.jpg
[picture2]: https://imgur.com/wj1ZgR5.jpg
[picture3]: https://imgur.com/NGOlxgY.jpg
[picture4]: https://imgur.com/HMmOBL0.jpg
[picture5]: https://imgur.com/H5k78vR.jpg
[picture6]: https://imgur.com/XTNxz7K.jpg
[picture7]: https://imgur.com/Pa9AEOB.jpg
[picture8]: https://imgur.com/ZID9EW1.jpg
[picture9]: https://imgur.com/BfAP5mB.jpg
[picture10]: https://imgur.com/8tGsDNW.jpg
[picture11]: https://imgur.com/5eCssZ4.jpg