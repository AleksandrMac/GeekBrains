package valid

import "strconv"

func isPort(port string) bool {
	n, _ := strconv.Atoi(port)
	return n < 65536
}

func isProtocol(protocol string) bool {
	switch protocol {
	case "git", "postgres", "http", "https", "ftp":
		return true
	default:
		return false
	}
}

func isURL(port string) bool {
	return false
}

/*ftp — протокол передачи файлов FTP
http — протокол передачи гипертекста HTTP
rtmp — проприетарный протокол потоковой передачи данных Real Time Messaging Protocol, в основном используется для передачи потокового видео и аудио с веб-камер через Интернет.
rtsp — потоковый протокол реального времени.
https — специальная реализация протокола HTTP, использующая шифрование (как правило, SSL или TLS)
gopher — протокол Gopher
mailto — адрес электронной почты
news — новости Usenet
nntp — новости Usenet через протокол NNTP
irc — протокол IRC
smb — протокол SMB/CIFS
prospero — служба каталогов Prospero Directory Service
telnet — ссылка на интерактивную сессию Telnet
wais — база данных системы WAIS
xmpp — протокол XMPP (часть Jabber)
file — имя локального файла
data — непосредственные данные (Data: URL)
tel — звонок по указанному телефону
Экзотические схемы URL:

afs — глобальное имя файла в файловой системе Andrew File System
cid — идентификатор содержимого для частей MIME
mid — идентификатор сообщений для электронной почты
mailserver — доступ к данным с почтовых серверов
nfs — имя файла в сетевой файловой системе NFS
tn3270 — эмуляция интерактивной сессии Telnet 3270
z39.50 — доступ к службам ANSI Z39.50
skype — протокол Skype
smsto — открытие редактора SMS в некоторых мобильных телефонах
ed2k — файлообменная сеть eDonkey, построенная по принципу P2P
market — Android-маркет
steam — протокол Steam
bitcoin — криптовалюта Биткойн
ob — OpenBazaar
tg — Telegram
Схемы URL в браузерах:

view-source — просмотр исходного кода указанной веб-страницы в различных браузерах.
В разных браузерах используются разные ключевые слова для доступа к служебным и сервисным страницам:
chrome в браузерах Google Chrome или браузеров на движке Gecko[4].
about в Firefox и других
opera в Opera
browser в Яндекс.Браузер.
*/
