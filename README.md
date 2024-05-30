Installation
---

Der Go-Installer kann von der Offiziellen Download-Page https://www.go.dev/doc/install heruntelgeladen werden.
Auf dieser Seite findet man auch die Installationsanleitungen für alle gängigen Betriebssysteme.
Nach der Installation kann man die Version über den Terminal mittels `go version` überprüfen.

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

Defer ist ein Schlüsselwort in Go, welches es ermöglicht, die Ausführung einer Funktion bis nach der Ausfühung der umgebenden Funktion zu verzögern. 
Auf den ersten Blick mag das nicht direkt sinnvoll erscheinen, da man die Funktion einfach am Ende aufrufen könnte. Defer bietet jedoch einige Vorteile, die besonders bei der Nutzung oder Analyse von Beispiel-Code erkennbar werden.

Ein häufiges Einsatzgebiet für defer ist das Lesen von Dateien. Oft vergisst man am Ende der Funktion den Inputstream zu schließen. Mit defer kann man die close()-Funktion direkt nach dem Öffnen der Datei aufrufen und muss sich später nicht mehr darum kümmern, den Stream zu schließen. Genau aus diesem Grund wird defer oft für Clean-up-Aufgaben verwendet.

Ein weiterer Vorteil von defer ist, dass es sicherstellt, dass eine Funktion ausgeführt wird, selbst im Falle eines Panics. Dadurch wird der Code übersichtlicher, da man nicht immer explizit kontrollieren muss, ob am Ende einer Funktion eine bestimmte Operation ausgeführt wird oder ob sie auch im Falle eines Panics ausgeführt wird.

Defer funktioniert nach dem LIFO-Prinzip (Last In, First Out), was bedeutet, dass das letzte defer-Statement als erstes ausgeführt wird. Durch dieses Prinzip kann man defer wie einen Stack verwenden.

Unterschiede zu Java
In Java wird das automatische Schließen von Ressourcen durch das Try-with-Resources-Konstrukt sichergestellt. Dieses Konstrukt stellt sicher, dass jede Ressource, die innerhalb des Try-Blocks deklariert wird, am Ende des Blocks automatisch geschlossen wird, selbst wenn eine Exception auftritt

Structural and Nominal typing
---

Nominal-Typisierung:

Nominal typing:
Nominal-Typisierung definiert Typen anhand ihrer Namen. Zwei Typen sind gleich, wenn ihre Namen übereinstimmen, unabhängig von ihrer Struktur.
Vorteile:
Klarheit und Ausdruckskraft: Typnamen geben eine klare und eindeutige Identität.
Sicherheit: Die Typen sind explizit, was die Codeverständlichkeit und Fehlererkennung verbessert.
Nachteile:
Weniger Flexibilität: Zwei Typen mit identischer Struktur, aber unterschiedlichen Namen werden als unterschiedliche Typen betrachtet, was die Wiederverwendbarkeit beeinträchtigen kann.
Boilerplate-Code: Bei ähnlichen Strukturen müssen separate Typen definiert werden, was zu wiederholtem Code führen kann.  
Strukturelle Typisierung:

Structural typing:
Strukturelle Typisierung definiert Typen anhand ihrer Struktur oder Form. Zwei Typen sind gleich, wenn ihre Strukturen übereinstimmen, unabhängig von ihren Namen.
Vorteile:
Flexibilität: Erlaubt die Verwendung von Typen mit ähnlichen Strukturen, ohne explizite Deklarationen.
Wiederverwendbarkeit: Ermöglicht den Einsatz von Typen, die nicht explizit für eine Schnittstelle definiert wurden.
Nachteile:
Verdeckung: Kann zu unerwartetem Verhalten führen, wenn Typen unbeabsichtigt dieselbe Struktur haben.
Vermischung von Konzepten: Die Identität eines Typs kann weniger klar sein, da sie von seiner Struktur und nicht von seinem Namen abhängt.

Go und Strukturelle Typisierung:
Go verwendet strukturelle Typisierung, insbesondere bei Schnittstellen. In Go kann ein Typ automatisch als eine Schnittstelle betrachtet werden, wenn er die erforderlichen Methoden implementiert, unabhängig vom Typnamen

In Java werdenTypen  durch ihre Namen definiert. Zwei Typen gelten als unterschiedlich, selbst wenn sie die gleiche Struktur haben, solange ihre Namen unterschiedlich sind. Dies führt zu mehr Sicherheit und Klarheit, kann aber die Wiederverwendbarkeit beeinträchtigen und mehr Boilerplate-Code erzeugen

Maps & Slices
---
### Maps ###

Maps:
Maps sind eine eingebaute Datenstruktur in Go, die Schlüssel-Wert-Paare speichert. Sie ermöglichen den schnellen Zugriff auf Werte über einen eindeutigen Schlüssel.
Vorteile:
Schneller Zugriff: Maps bieten eine schnelle Suche nach Werten anhand ihres Schlüssels.
Dynamische Größe: Maps sind dynamisch dimensioniert und können Elemente hinzufügen oder entfernen, während das Programm läuft.
Nachteile:
Ungeordnet: Die Reihenfolge der Elemente in einer Map ist nicht garantiert und kann variieren.
Speicherbedarf: Maps benötigen mehr Speicher als Arrays oder Slices, was bei großen Datenmengen zu höherem Speicherverbrauch führen kann.
### Slices ###

Slices:
Slices sind eine Ansicht von Arrays in Go, die eine flexible, dynamische Größe ermöglichen. Sie können Elemente hinzufügen, entfernen oder unterteilen.
Vorteile:
Flexibel: Slices können leicht verändert werden, indem Elemente hinzugefügt, entfernt oder unterteilt werden.
Effizienter Speicher: Slices nutzen weniger Speicher als Arrays und ermöglichen eine effiziente Verwaltung von Daten.
Nachteile:
Verkettete Struktur: Die interne Struktur von Slices als Ansicht von Arrays kann zu Laufzeitkosten führen, insbesondere bei großen Datenmengen.
Dynamische Größe: Die dynamische Größe von Slices kann zu häufigen Allokationen und Freigaben führen, was die Leistung beeinträchtigen kann.

Java bietet ähnliche Funktionalitäten wie Go, jedoch mit expliziten Klassen und Methoden. Go bietet eine kompaktere und weniger boilerplate-lastige Syntax, insbesondere bei der Handhabung von dynamischen Datenstrukturen.
In Java entsprechen HashMaps den Maps in Go, und ArrayLists den Slices.

Goroutine  
---
Goroutinen sind leichtgewichtige Threads, die von der Go-Runtime verwaltet werden. Im Gegensatz zu herkömmlichen Threads sind Goroutinen nicht an Betriebssystemthreads gebunden, was bedeutet, dass Tausende von Goroutinen auf einem einzigen Betriebssystemthread laufen können. Dies ermöglicht es Go, große Mengen von gleichzeitigen Aufgaben effizient zu verarbeiten
Goroutinen werden durch das Schlüsselwort "go" gestartet. Wenn Sie eine Funktion mit "go" aufrufen, wird sie als Goroutine ausgeführt. Goroutinen teilen sich den Adressraum derselben Anwendung und kommunizieren über Kanäle oder durch gemeinsam genutzte Speicherbereiche.  

Nebenläufigkeit in Java
Java nutzt traditionelle Threads zur Handhabung von Nebenläufigkeit. Ein Thread ist eine unabhängige Ausführungseinheit innerhalb eines Prozesses, die es ermöglicht, mehrere Aufgaben gleichzeitig auszuführen. Java bietet zwei Hauptmechanismen zur Implementierung von Threads: durch Erben von der Thread-Klasse oder durch Implementieren des Runnable-Interfaces.
Zusätzlich bietet Java das Executor-Framework zur Verwaltung von Threads und Aufgaben. Das Framework vereinfacht die Handhabung von Thread-Pools und die Zuweisung von Aufgaben, was eine effizientere Thread-Nutzung gewährleistet.

Coroutinen
Coroutine ist ein allgemeines Konzept in der Informatik, das eine Unterbrechung der Ausführung eines Programms an einer bestimmten Stelle und die Fortsetzung an einer anderen Stelle ermöglicht. Im Gegensatz zu Threads sind Coroutinen kooperativ und teilen sich den gleichen Speicherbereich. Sie werden nicht vom Betriebssystem verwaltet, sondern von der Laufzeitumgebung der Programmiersprache.

GoRoutinen
Goroutinen sind spezifisch für die Go-Programmiersprache und stellen eine Implementierung von Coroutinen dar. Sie sind leichtgewichtige Threads, die von der Go-Runtime verwaltet werden. Im Gegensatz zu herkömmlichen Threads sind Goroutinen nicht an Betriebssystemthreads gebunden, was bedeutet, dass Tausende von Goroutinen auf einem einzigen Betriebssystemthread laufen können. Dies ermöglicht es Go, große Mengen von gleichzeitigen Aufgaben effizient zu verarbeiten.

Unterschied zwischen go- & Coroutinen
Der Hauptunterschied zwischen Goroutinen und Coroutinen liegt in ihrer Implementierung und Verwaltung. Während Goroutinen spezifisch für die Go-Sprache sind und von der Go-Runtime verwaltet werden, sind Coroutinen ein allgemeines Konzept, das in verschiedenen Programmiersprachen implementiert sein kann, aber möglicherweise unterschiedliche Eigenschaften und Verhaltensweisen aufweist. Goroutinen sind daher eine spezifische Implementierung von Coroutinen in Go, die auf die Bedürfnisse und Eigenschaften der Sprache zugeschnitten sind.

Channels & Select  
---
### Channels ###
Channels ist eine Konstruktionsform, die es Goroutinen ermöglicht, sicher Daten miteinander auszutauschen und zu synchronisieren, ohne dass explizite Sperren oder Mutexe verwendet werden müssen. Ein Kanal ist ein typsicherer Datenstrom, der Daten zwischen Goroutinen überträgt. Kanäle sind asynchron, was bedeutet, dass Sender und Empfänger nicht gleichzeitig bereit sein müssen, um eine Nachricht zu senden oder zu empfangen. Wenn ein Sender eine Nachricht in einen Kanal schreibt, blockiert die Operation, bis ein Empfänger die Nachricht liest. Ebenso blockiert das Lesen aus einem Kanal, bis eine Nachricht verfügbar ist

### Select ###
Das select-Statement in Go wird verwendet, um auf mehrere Kanäle gleichzeitig zu warten und die erste Operation auszuführen, die bereit ist. Es ist ähnlich wie ein Switch-Statement, aber speziell für Kanäle. Wenn nicht klar ist, welcher von mehreren Channel zuerst verfügbar sein wird (Daten enthält) kann Select verwendet werden, um auf das zu reagieren, was als nächstes verfügbar ist. Dies ist besonders nützlich, wenn Sie auf eingehende Daten von verschiedenen Goroutinen oder Quellen warten, und Sie möchten nicht blockieren, bis alle Kanäle Daten senden.

Anstatt von Channels und Select verwendet Java ähnliche Konzepte wie die BlockingQueue und ExecutorService, um die Kommunikation und Synchronisation zwischen Threads zu ermöglichen. Die BlockingQueue fungiert als ein Äquivalent zu Kanälen in Go und ermöglicht es Threads, Daten sicher auszutauschen und zu synchronisieren. Der ExecutorService ermöglicht es, Aufgaben an Threads zu übergeben und das Ergebnis zurückzugeben, ähnlich wie das Select-Statement in Go, das auf mehrere Kanäle wartet und die erste verfügbare Operation ausführt

Package Management
---
In Go wird das Paketmanagement durch das Tool "go mod" verwaltet. Mit go mod können Entwickler Abhängigkeiten zu externen Paketen verwalten. Dabei werden die benötigten Pakete und deren Versionen in einer "go.mod"-Datei aufgeführt. Go mod stellt sicher, dass die richtigen Versionen der Pakete heruntergeladen und in das Projekt integriert werden. Externe Pakete werden standardmäßig im GOPATH des Benutzers oder im Modulcache von Go gespeichert.

In Java wird das Paketmanagement oft über Build-Management-Tools wie Maven oder Gradle verwaltet. Diese Tools ermöglichen es, Abhängigkeiten zu externen Bibliotheken in einem Projekt zu definieren, und sie kümmern sich um das Herunterladen und Verwalten dieser Abhängigkeiten. Die Abhängigkeiten und ihre Versionen werden in einer Konfigurationsdatei (z. B. "pom.xml" für Maven oder "build.gradle" für Gradle) aufgeführt. Die benötigten Bibliotheken werden von einem zentralen Repository wie Maven Central heruntergeladen und in das Projekt integriert.
