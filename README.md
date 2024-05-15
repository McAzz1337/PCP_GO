Installation
---

Der Go-Installer kann von der Offiziellen Download-Page https://www.go.dev/doc/install heruntelgeladen werden.
Auf dieser Seite findet man auch die Installationsanleitungen für alle gängigen Betriebssysteme.
Nach der Installation kann man die Version über den Terminal mittels: go version

Go code schreiben
---
Go-Code kann in einem beliebigen Texteditor geschrieben werden. Mit einer IDE oder einem Codeeditor ist die
Entwicklung jedooch einiges einfacher, da man einen Linter, Code-Completion, Syntax-Highlighting und weitere
Features zur Verfügung und Unterstützung hat.
Mit dem Codeeditor VsCode fällt es einfach und dauert nicht lange, um loszulegen mit dem Programmieren. Dazu
muss man nur die offizielle Extension "Go" von Go Team at Google zu installieren.

Projekt
---

Ein go Projekt kann auf sehr einfache Art und weise Erstellt werden. Dazu muss man lediglich in einem beliebigen
Verzeichniss eine Datei mit der Endung .go erstellen(main.go muss den Einstiegspunkt "func main(){}" enthalten).
Um den Code auszuführen benutzt man den Befehl "go run <file>.go". Jedoch wird nur diese Datei dann ausgeführt,
und es kommt zu Fehlern, wenn ein Symbol aus einer anderen Datei referenziert wird. Um mehrere Dateien ausführen
zu können, kann man den Befehl "go run <fil1>.go <fil2>.go..." oder "go run *.go" benutzen.
Mit diesem Befehl wird aber der code nur ausgeführt, aber nicht zu einer executable kompiliert.
Um eine executable zu erstellen braucht man in einem Project ein Modul benötigt. Dies kann mithilfe des Befehls
"go mod init <module_name>" estellt werden. Danach kann man mit "go build" die executable erstelllen,
und mit "./<executable_name> ausgeführt weden, wobei executable_name der selbe name wie module_name ist.


Defer
---

Defer ist ein Schlüsselwort, welches es ermoglicht, die Ausführung einer Funktion an das Ende einer Funktion zu verschieben.
Macht nicht direkt Sinn, da man die Funktion einfach am Ende Aufrufen kann, hat aber ein paar Vorteile, die erkennbar werden,
wenn man es selbst benutzt, oder Beispiel-Code analysiert.
Defer ist sehr nützlch zum Beispiel für Operationen wie lesen einer Datei. Oft vergisst man am Ende der Funktion den Inputstream
wieder zu schliessen. Man kann die "close()" Funktion mit defer aufrufen direkt nadem man eine Referenz zu der Datei hat und muss
sich nacher nicht mehr darum kümmer diese zu schliessen. Genau darum wird defer oft für clean-up Aufgaben verwendet.
Defer hat auch den Vortril, dass es sicherstellt, dass eine Funktion ausgeführt wird, selbst im falle eines panics.
Der Code wird dadurch auch übersichtlicher, da man nicht immer kontrollieren muss, ob am Ende einer Funktion etwas bestimmtes,
passiert, oder im Falle eines panics, diese auch noch ausgeführt wird.
Defer Funktioniert nach den LIFO-Prinzip, darum wird das letzte defer Statement als erstes ausgeführt. Durch dieses Prinzip kann man
damit wie mit einem Stack umgehen.

Structural and Nominal typing
---

Nominal-Typisierung:

Erklärung:
Nominal-Typisierung definiert Typen anhand ihrer Namen. Zwei Typen sind gleich, wenn ihre Namen übereinstimmen, unabhängig von ihrer Struktur.
Vorteile:
Klarheit und Ausdruckskraft: Typnamen geben eine klare und eindeutige Identität.
Sicherheit: Die Typen sind explizit, was die Codeverständlichkeit und Fehlererkennung verbessert.
Nachteile:
Weniger Flexibilität: Zwei Typen mit identischer Struktur, aber unterschiedlichen Namen werden als unterschiedliche Typen betrachtet, was die Wiederverwendbarkeit beeinträchtigen kann.
Boilerplate-Code: Bei ähnlichen Strukturen müssen separate Typen definiert werden, was zu wiederholtem Code führen kann.  
Strukturelle Typisierung:

Erklärung:
Strukturelle Typisierung definiert Typen anhand ihrer Struktur oder Form. Zwei Typen sind gleich, wenn ihre Strukturen übereinstimmen, unabhängig von ihren Namen.
Vorteile:
Flexibilität: Erlaubt die Verwendung von Typen mit ähnlichen Strukturen, ohne explizite Deklarationen.
Wiederverwendbarkeit: Ermöglicht den Einsatz von Typen, die nicht explizit für eine Schnittstelle definiert wurden.
Nachteile:
Verdeckung: Kann zu unerwartetem Verhalten führen, wenn Typen unbeabsichtigt dieselbe Struktur haben.
Vermischung von Konzepten: Die Identität eines Typs kann weniger klar sein, da sie von seiner Struktur und nicht von seinem Namen abhängt.

Maps & Slices
---
Maps:

Erklärung:
Maps sind eine eingebaute Datenstruktur in Go, die Schlüssel-Wert-Paare speichert. Sie ermöglichen den schnellen Zugriff auf Werte über einen eindeutigen Schlüssel.
Vorteile:
Schneller Zugriff: Maps bieten eine schnelle Suche nach Werten anhand ihres Schlüssels.
Dynamische Größe: Maps sind dynamisch dimensioniert und können Elemente hinzufügen oder entfernen, während das Programm läuft.
Nachteile:
Ungeordnet: Die Reihenfolge der Elemente in einer Map ist nicht garantiert und kann variieren.
Speicherbedarf: Maps benötigen mehr Speicher als Arrays oder Slices, was bei großen Datenmengen zu höherem Speicherverbrauch führen kann.
Slices:

Erklärung:
Slices sind eine Ansicht von Arrays in Go, die eine flexible, dynamische Größe ermöglichen. Sie können Elemente hinzufügen, entfernen oder unterteilen.
Vorteile:
Flexibel: Slices können leicht verändert werden, indem Elemente hinzugefügt, entfernt oder unterteilt werden.
Effizienter Speicher: Slices nutzen weniger Speicher als Arrays und ermöglichen eine effiziente Verwaltung von Daten.
Nachteile:
Verkettete Struktur: Die interne Struktur von Slices als Ansicht von Arrays kann zu Laufzeitkosten führen, insbesondere bei großen Datenmengen.
Dynamische Größe: Die dynamische Größe von Slices kann zu häufigen Allokationen und Freigaben führen, was die Leistung beeinträchtigen kann.

Goroutine  
---
tbd

Channels & Select  
---
tbd

Package Management
---
tbd