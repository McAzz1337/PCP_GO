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

