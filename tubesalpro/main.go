
package main

import (
	"net/http"
  "html/template"
  "log"
  "math/rand"
  "strconv"
  "time"
)

var tmpl, tmpl2, tmpl3, tmpl4, tmpl5, tmpl6, tmpl7, tmpl8, tmpl9, tmpl10, tmpl11, tmpl12,tmpl13, tmpl14, tmpl15 *template.Template
const maxItems = 100
var players Peserta
var items Question

type Peserta struct {
  arrPlayers [maxItems]Player
  playerCount int
}

type Question struct {
  arrItems [maxItems]Item 
  itemCount int
}

type Item struct {
  Name string
  Jwba string
  Jwbb string 
  Jwbc string 
  Jwbd string
  Kunci string
  Unique string
  PCorrect int
  PFalse int
  Index int
}

type PageData struct {
  Title string
}

type Player struct {
  Name string 
  Correct int
  Result string
}

func menu(w http.ResponseWriter, r *http.Request) {
  data := PageData{
    Title: "Who One to Be a Millionaire!",
  }
  
  tmpl.Execute(w, data)
}

func tambahForm(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    if items.itemCount < maxItems && r.FormValue("soal") != ""{
      name := r.FormValue("soal")
      items.arrItems[items.itemCount].Name = name
      jwba := r.FormValue("jwba")
      items.arrItems[items.itemCount].Jwba = jwba
      jwbb := r.FormValue("jwbb")
      items.arrItems[items.itemCount].Jwbb = jwbb
      jwbc := r.FormValue("jwbc")
      items.arrItems[items.itemCount].Jwbc = jwbc
      jwbd := r.FormValue("jwbd")
      items.arrItems[items.itemCount].Jwbd = jwbd
      kunci := r.FormValue("kunci")
      items.arrItems[items.itemCount].Kunci = kunci
      unique := r.FormValue("unique")
      items.arrItems[items.itemCount].Unique = unique
      items.itemCount++
    }
  }

  err := tmpl2.Execute(w, items.arrItems[:items.itemCount])
  if err != nil {
    log.Println(err)
  }
}

func ubahForm(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    key := r.FormValue("unique")
    for i := 0 ; i < items.itemCount; i++ {
      if items.arrItems[i].Unique == key {
        soalBaru := r.FormValue("new")
        items.arrItems[i].Name = soalBaru
        jwbaBaru := r.FormValue("jwbaBaru")
        items.arrItems[i].Jwba = jwbaBaru
        jwbbBaru := r.FormValue("jwbbBaru")
        items.arrItems[i].Jwbb = jwbbBaru
        jwbcBaru := r.FormValue("jwbcBaru")
        items.arrItems[i].Jwbc = jwbcBaru
        jwbdBaru := r.FormValue("jwbdBaru")
        items.arrItems[i].Jwbd = jwbdBaru
        kunciBaru := r.FormValue("kunciBaru")
        items.arrItems[i].Kunci = kunciBaru
        newUnique := r.FormValue("uniqueBaru")
        items.arrItems[i].Unique = newUnique
      }
    }
  }

  err := tmpl4.Execute(w, items.arrItems[:items.itemCount])
  if err != nil {
    log.Println(err)
  }  
}

func deleteForm(w http.ResponseWriter, r *http.Request) {
  var index int
  if r.Method == "POST" {
    key := r.FormValue("unique")
    for j := 0; j < items.itemCount; j++ {
      if items.arrItems[j].Unique == key {
        index = j
      }
    }
    for i := index; i < items.itemCount-1; i++ {
      items.arrItems[i].Name = items.arrItems[i+1].Name
      items.arrItems[i].Jwba = items.arrItems[i+1].Jwba
      items.arrItems[i].Jwbb = items.arrItems[i+1].Jwbb
      items.arrItems[i].Jwbc = items.arrItems[i+1].Jwbc
      items.arrItems[i].Jwbd = items.arrItems[i+1].Jwbd
      items.arrItems[i].Kunci = items.arrItems[i+1].Kunci
      items.arrItems[i].Unique = items.arrItems[i+1].Unique
    }
    items.arrItems[items.itemCount-1].Name = ""
    items.arrItems[items.itemCount-1].Jwba = ""
    items.arrItems[items.itemCount-1].Jwbb = ""
    items.arrItems[items.itemCount-1].Jwbc = ""
    items.arrItems[items.itemCount-1].Jwbd = ""
    items.arrItems[items.itemCount-1].Kunci = ""
    items.arrItems[items.itemCount-1].Unique = ""
    items.itemCount--
  }

  err := tmpl5.Execute(w, items.arrItems[:items.itemCount])
  if err != nil {
    log.Println(err)
  }  
}

func admin(w http.ResponseWriter, r *http.Request) {
  err := tmpl3.Execute(w, items.arrItems[:items.itemCount])
  if err != nil {
    log.Println(err)
  }
}


func leaderboard(w http.ResponseWriter, r *http.Request) {
  var pass, i, idx, temp int
  var temp2 string
  pass = 1
    for pass <= players.playerCount-1 {
      idx = pass -1 
      i = pass 
      for i < players.playerCount {
        if players.arrPlayers[idx].Correct < players.arrPlayers[i].Correct {
          idx = i
        }
        i++
      }
      temp = players.arrPlayers[pass-1].Correct 
      temp2 = players.arrPlayers[pass-1].Name
      players.arrPlayers[pass-1].Correct = players.arrPlayers[idx].Correct
      players.arrPlayers[pass-1].Name = players.arrPlayers[idx].Name
      players.arrPlayers[idx].Correct = temp
      players.arrPlayers[idx].Name = temp2
      pass++
    } 

  err := tmpl13.Execute(w, players.arrPlayers[:players.playerCount])
  if err != nil {
    log.Println(err)
  }
}

func benar(w http.ResponseWriter, r *http.Request) {
  var copyBenar [maxItems]Item 
  var pass, i, idx, temp int
  var temp2 string
  copyBenar = items.arrItems
  pass = 1
    for pass <= 5 {
      idx = pass -1 
      i = pass 
      for i < items.itemCount {
        if copyBenar[idx].PCorrect < copyBenar[i].PCorrect {
          idx = i
        }
        i++
      }
      temp = copyBenar[pass-1].PCorrect 
      temp2 = copyBenar[pass-1].Name
      copyBenar[pass-1].PCorrect = copyBenar[idx].PCorrect
      copyBenar[pass-1].Name = copyBenar[idx].Name
      copyBenar[idx].PCorrect = temp
      copyBenar[idx].Name = temp2
      pass++
    } 

    err := tmpl14.Execute(w, copyBenar[:5])
  if err != nil {
    log.Println(err)
  }
}

func salah(w http.ResponseWriter, r *http.Request) {
  var copySalah [maxItems]Item 
  var pass, i, idx, temp int
  var temp2 string
  copySalah = items.arrItems
  pass = 1
    for pass <= 5 {
      idx = pass -1 
      i = pass 
      for i < items.itemCount {
        if copySalah[idx].PFalse < copySalah[i].PFalse {
          idx = i
        }
        i++
      }
      temp = copySalah[pass-1].PFalse
      temp2 = copySalah[pass-1].Name
      copySalah[pass-1].PFalse = copySalah[idx].PFalse
      copySalah[pass-1].Name = copySalah[idx].Name
      copySalah[idx].PFalse = temp
      copySalah[idx].Name = temp2
      pass++
    } 

    err := tmpl15.Execute(w, copySalah[:5])
  if err != nil {
    log.Println(err)
  }
}

func main() {
  var currentPIndex int

  // Default Soal & Jawaban
  items.itemCount = 20
  items.arrItems[0].Name = "Apa nama ibukota Australia?"
  items.arrItems[1].Name = "Siapa Presiden Indonesia saat ini?"
  items.arrItems[2].Name = "Ada berapa provinsi di indonesia saat ini?"
  items.arrItems[3].Name = "Apa nama buah yang dimakan Nabi Adam sehingga diturunkan ke Bumi?"
  items.arrItems[4].Name = "Apa nama ibukota Indonesia?"
  items.arrItems[5].Name = "Tumbuhan berduri yang kebanyakan tumbuh pada daerah gersang seperti gurun adalah?"
  items.arrItems[6].Name = "Negara terluas di dunia adalah?"
  items.arrItems[7].Name = "Tari kecak adalah tari yang berasal dari daerah? "
  items.arrItems[8].Name = "Gunung tertinggi di dunia adalah?"
  items.arrItems[9].Name = "Mata uang negara Jepang yaitu?"
  items.arrItems[10].Name = "Apa ibu kota Portugal?"
  items.arrItems[11].Name = "Bahan bakar kereta api adalah?"
  items.arrItems[12].Name = "Vitamin yang banyak terkandung dalam buah-buahan adalah?"
  items.arrItems[13].Name = "Pusat keuangan kota Amerika Serikat adalah?"
  items.arrItems[14].Name = "Kota paling boros listrik di Asia adalah?"
  items.arrItems[15].Name = "Burung tercepat di dunia adalah?"
  items.arrItems[16].Name = "Negara terkaya di dunia adalah?"
  items.arrItems[17].Name = "Patung Liberty dibuat di negara?"
  items.arrItems[18].Name = "Minuman terfavorit di dunia adalah?"
  items.arrItems[19].Name = "Jumlah warna pelangi ada?"


  items.arrItems[0].Jwba = "Sydney"
  items.arrItems[0].Jwbb = "Canberra"
  items.arrItems[0].Jwbc = "Jakarta"
  items.arrItems[0].Jwbd = "Melbourne"
  items.arrItems[0].Kunci = "Canberra"
  items.arrItems[0].Unique = "AusCap"

  items.arrItems[1].Jwba = "Donald Trump"
  items.arrItems[1].Jwbb = "Megawati"
  items.arrItems[1].Jwbc = "Puan Maharani"
  items.arrItems[1].Jwbd = "Joko Widodo"
  items.arrItems[1].Kunci = "Joko Widodo"
  items.arrItems[1].Unique = "IndoPres"

  items.arrItems[2].Jwba = "39"
  items.arrItems[2].Jwbb = "38"
  items.arrItems[2].Jwbc = "40"
  items.arrItems[2].Jwbd = "27"
  items.arrItems[2].Kunci = "38"
  items.arrItems[2].Unique = "IndoProv"

  items.arrItems[3].Jwba = "Apel"
  items.arrItems[3].Jwbb = "Pisang"
  items.arrItems[3].Jwbc = "Khuldi"
  items.arrItems[3].Jwbd = "Manggis"
  items.arrItems[3].Kunci = "Khuldi"
  items.arrItems[3].Unique = "AdamFruit"

  items.arrItems[4].Jwba = "Jakarta"
  items.arrItems[4].Jwbb = "Canberra"
  items.arrItems[4].Jwbc = "Tokyo"
  items.arrItems[4].Jwbd = "Melbourne"
  items.arrItems[4].Kunci = "Jakarta"
  items.arrItems[4].Unique = "IndoCap"

  items.arrItems[5].Jwba = "Kaktus"
  items.arrItems[5].Jwbb = "Bunga Matahari"
  items.arrItems[5].Jwbc = "Pohon Cemara"
  items.arrItems[5].Jwbd = "Mawar"
  items.arrItems[5].Kunci = "Kaktus"
  items.arrItems[5].Unique = "TumbuhanGurun"

  items.arrItems[6].Jwba = "Amerika"
  items.arrItems[6].Jwbb = "Indonesia"
  items.arrItems[6].Jwbc = "China"
  items.arrItems[6].Jwbd = "Rusia"
  items.arrItems[6].Kunci = "Rusia"
  items.arrItems[6].Unique = "NegaraTerluas"

	items.arrItems[7].Jwba = "Bali"
  items.arrItems[7].Jwbb = "Sunda"
	items.arrItems[7].Jwbc = "Aceh"
	items.arrItems[7].Jwbd = "Minang"
	items.arrItems[7].Kunci = "Bali"
	items.arrItems[7].Unique ="QuestTari"

  items.arrItems[8].Jwba = "Gunung Everest"
  items.arrItems[8].Jwbb = "Gunung Tangkuban Perahu"
  items.arrItems[8].Jwbc = "Gunung Aconcagua"
  items.arrItems[8].Jwbd = "Gunung McKinley"
  items.arrItems[8].Kunci = "Gunung Everest"
  items.arrItems[8].Unique = "GunungTertinggiDidunia"

  items.arrItems[9].Jwba = "Yen"
  items.arrItems[9].Jwbb = "Rupiah"
  items.arrItems[9].Jwbc = "Ringgit"
  items.arrItems[9].Jwbd = "Dolar"
  items.arrItems[9].Kunci = "Yen"
  items.arrItems[9].Unique = "MataUangJepang"

  items.arrItems[10].Jwba = "Lisbon"
  items.arrItems[10].Jwbb = "Kuala Lumpur"
  items.arrItems[10].Jwbc = "Jakarta"
  items.arrItems[10].Jwbd = "Bangkok"
  items.arrItems[10].Kunci = "Lisbon"
	items.arrItems[10].Unique ="QuestIbukotaPort"

  items.arrItems[11].Jwba = "Batu bara"
	items.arrItems[11].Jwbb = "Bensin"
	items.arrItems[11].Jwbc = "Minyak tanah"
	items.arrItems[11].Jwbd = "Minyak solar"
	items.arrItems[11].Kunci = "Batu bara"
	items.arrItems[11].Unique ="QuestKerapi"

  items.arrItems[12].Jwba = "Vitamin A"
  items.arrItems[12].Jwbb = "Vitamin B"
	items.arrItems[12].Jwbc = "Vitamin C"
	items.arrItems[12].Jwbd = "Vitamin D"
	items.arrItems[12].Kunci = "Vitamin C"
	items.arrItems[12].Unique ="QuestVit"

  items.arrItems[13].Jwba = "New York"
	items.arrItems[13].Jwbb = "Chicago"
	items.arrItems[13].Jwbc = "Los Angeles"
	items.arrItems[13].Jwbd = "San Francisco"
	items.arrItems[13].Kunci = "New York"
	items.arrItems[13].Unique ="QuestAmrik"

  items.arrItems[14].Jwba = "Tokyo"
  items.arrItems[14].Jwbb = "Osaka"
  items.arrItems[14].Jwbc = "Jakarta"
  items.arrItems[14].Jwbd = "Bangkok"
  items.arrItems[14].Kunci = "Tokyo"
  items.arrItems[14].Unique ="QuestAsia"

  items.arrItems[15].Jwba = "Falcon"
	items.arrItems[15].Jwbb = "Burung Unta"
	items.arrItems[15].Jwbc = "Burung Hantu"
	items.arrItems[15].Jwbd = "Elang"
	items.arrItems[15].Kunci = "Falcon"
	items.arrItems[15].Unique ="QuestBird"

  items.arrItems[16].Jwba = "Qatar"
  items.arrItems[16].Jwbb = "Amerika"
  items.arrItems[16].Jwbc = "Korea Selatan"
  items.arrItems[16].Jwbd = "Indonesia"
  items.arrItems[16].Kunci = "Qatar"
  items.arrItems[16].Unique ="QuestKaya"

  items.arrItems[17].Jwba = "Indonesia"
	items.arrItems[17].Jwbb = "Prancis"
	items.arrItems[17].Jwbc = "Amerika"
	items.arrItems[17].Jwbd = "Italia"
	items.arrItems[17].Kunci = "Prancis"
	items.arrItems[17].Unique ="QuestLiberty"

  items.arrItems[18].Jwba = "Teh"
  items.arrItems[18].Jwbb = "Jus"
  items.arrItems[18].Jwbc = "Kopi"
  items.arrItems[18].Jwbd = "Es Cendol"
  items.arrItems[18].Kunci = "Teh"
  items.arrItems[18].Unique ="QuestDrink"

  items.arrItems[19].Jwba = "7"
	items.arrItems[19].Jwbb = "8"
	items.arrItems[19].Jwbc = "9"
	items.arrItems[19].Jwbd = "10"
	items.arrItems[19].Kunci = "7"
	items.arrItems[19].Unique ="QuestPelangi"
  

  mux := http.NewServeMux()
  tmpl = template.Must(template.ParseFiles("index.html"))
  tmpl2 = template.Must(template.ParseFiles("add.html"))
  tmpl3 = template.Must(template.ParseFiles("admin.html"))
  tmpl4 = template.Must(template.ParseFiles("edit.html"))
  tmpl5 = template.Must(template.ParseFiles("delete.html"))
  tmpl6 = template.Must(template.ParseFiles("peserta.html"))
  tmpl7 = template.Must(template.ParseFiles("play1.html"))
  tmpl8 = template.Must(template.ParseFiles("play2.html")) 
  tmpl9 = template.Must(template.ParseFiles("play3.html")) 
  tmpl10 = template.Must(template.ParseFiles("play4.html")) 
  tmpl11 = template.Must(template.ParseFiles("play5.html")) 
  tmpl12 = template.Must(template.ParseFiles("result.html")) 
  tmpl13 = template.Must(template.ParseFiles("leaderboard.html")) 
  tmpl14 = template.Must(template.ParseFiles("benar.html")) 
  tmpl15 = template.Must(template.ParseFiles("salah.html")) 


  mux.HandleFunc("/", menu)
  mux.HandleFunc("/admin", admin)
  mux.HandleFunc("/edit", ubahForm)
  mux.HandleFunc("/add", tambahForm)
  mux.HandleFunc("/hapus", deleteForm)
  mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
    if r.FormValue("submit") == "saveAndRedirect" {
      playerName := r.FormValue("player")
      var i, ada int 
      ada = -1
      for i = 0; i < players.playerCount; i ++ {
        if playerName == players.arrPlayers[i].Name {
          ada = i
        }
      }
      if ada == -1 {
        players.arrPlayers[players.playerCount].Name = playerName
        players.playerCount+= 1 
        currentPIndex = players.playerCount - 1
      }else {
        players.arrPlayers[ada].Correct = 0
        currentPIndex = ada
      }
      
      http.Redirect(w, r, "/soal1", http.StatusFound)
    }
  
    err := tmpl6.Execute(w, players.arrPlayers[:players.playerCount])
    if err != nil {
      log.Println(err)
    }
  })
  mux.HandleFunc("/soal1", func(w http.ResponseWriter, r *http.Request) {
    rand.Seed(time.Now().UnixNano())
    index := rand.Intn(items.itemCount)
  
    items.arrItems[index].Index = index
  
    if r.FormValue("submit") == "saveAndRedirect" {
      answer1 := r.FormValue("jwb1")
      index := r.FormValue("index")
      i, irr := strconv.Atoi(index)
      if irr != nil {
        http.Error(w, "Invalid random index", http.StatusBadRequest)
        return
      }
      if answer1 == items.arrItems[i].Kunci {
        players.arrPlayers[currentPIndex].Correct += 1
        items.arrItems[i].PCorrect += 1
      }else if answer1 != items.arrItems[i].Kunci && answer1 != ""{
        items.arrItems[i].PFalse += 1
      }
      http.Redirect(w, r, "/soal2", http.StatusFound)
    }
  
    err := tmpl7.Execute(w, items.arrItems[items.arrItems[index].Index])
    if err != nil {
      log.Println(err)
    }
  })
  mux.HandleFunc("/soal2", func(w http.ResponseWriter, r *http.Request) {
    rand.Seed(time.Now().UnixNano())
    index := rand.Intn(items.itemCount)
  
    items.arrItems[index].Index = index
  
    if r.FormValue("submit") == "saveAndRedirect" {
      answer2 := r.FormValue("jwb2")
      index := r.FormValue("index")
      i, irr := strconv.Atoi(index)
      if irr != nil {
        http.Error(w, "Invalid random index", http.StatusBadRequest)
        return
      }
      if answer2 == items.arrItems[i].Kunci {
        players.arrPlayers[currentPIndex].Correct += 1
        items.arrItems[i].PCorrect += 1
      }else if answer2 != items.arrItems[i].Kunci && answer2 != ""{
        items.arrItems[i].PFalse += 1
      }
      http.Redirect(w, r, "/soal3", http.StatusFound)
    }
  
    err := tmpl8.Execute(w, items.arrItems[items.arrItems[index].Index])
    if err != nil {
      log.Println(err)
    }
  })
  mux.HandleFunc("/soal3", func(w http.ResponseWriter, r *http.Request) {
    rand.Seed(time.Now().UnixNano())
    index := rand.Intn(items.itemCount)
  
    items.arrItems[index].Index = index
  
    if r.FormValue("submit") == "saveAndRedirect" {
      answer3 := r.FormValue("jwb3")
      index := r.FormValue("index")
      i, irr := strconv.Atoi(index)
      if irr != nil {
        http.Error(w, "Invalid random index", http.StatusBadRequest)
        return
      }
      if answer3 == items.arrItems[i].Kunci {
        players.arrPlayers[currentPIndex].Correct += 1
        items.arrItems[i].PCorrect += 1
      }else if answer3 != items.arrItems[i].Kunci && answer3 != ""{
        items.arrItems[i].PFalse += 1
      }
      http.Redirect(w, r, "/soal4", http.StatusFound)
    }
  
    err := tmpl9.Execute(w, items.arrItems[items.arrItems[index].Index])
    if err != nil {
      log.Println(err)
    }
  })
  mux.HandleFunc("/soal4", func(w http.ResponseWriter, r *http.Request) {
    rand.Seed(time.Now().UnixNano())
    index := rand.Intn(items.itemCount)
  
    items.arrItems[index].Index = index
  
    if r.FormValue("submit") == "saveAndRedirect" {
      answer4 := r.FormValue("jwb4")
      index := r.FormValue("index")
      i, irr := strconv.Atoi(index)
      if irr != nil {
        http.Error(w, "Invalid random index", http.StatusBadRequest)
        return
      }
      if answer4 == items.arrItems[i].Kunci {
        players.arrPlayers[currentPIndex].Correct += 1
        items.arrItems[i].PCorrect += 1
      }else if answer4 != items.arrItems[i].Kunci && answer4 != ""{
        items.arrItems[i].PFalse += 1
      }
      http.Redirect(w, r, "/soal5", http.StatusFound)
    }
  
    err := tmpl10.Execute(w, items.arrItems[items.arrItems[index].Index])
    if err != nil {
      log.Println(err)
    }
  })
  mux.HandleFunc("/soal5", func(w http.ResponseWriter, r *http.Request) {
    rand.Seed(time.Now().UnixNano())
    index := rand.Intn(items.itemCount)
  
    items.arrItems[index].Index = index
  
    if r.FormValue("submit") == "saveAndRedirect" {
      answer5 := r.FormValue("jwb5")
      index := r.FormValue("index")
      i, irr := strconv.Atoi(index)
      if irr != nil {
        http.Error(w, "Invalid random index", http.StatusBadRequest)
        return
      }
      if answer5 == items.arrItems[i].Kunci {
        players.arrPlayers[currentPIndex].Correct += 1
        items.arrItems[i].PCorrect += 1
      }else if answer5 != items.arrItems[i].Kunci && answer5 != ""{
        items.arrItems[i].PFalse += 1
      }
      http.Redirect(w, r, "/result", http.StatusFound)
    }
  
  
    err := tmpl11.Execute(w, items.arrItems[items.arrItems[index].Index])
    if err != nil {
      log.Println(err)
    }
  })
  mux.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
    if players.arrPlayers[currentPIndex].Correct == 5 {
      players.arrPlayers[currentPIndex].Result = "Selamat kamu berhasil menjadi Millionaire!"
    }else {
      players.arrPlayers[currentPIndex].Result = "Maaf kamu gagal menjadi Millionaire :( silahkan coba lagi!"
    }
    err := tmpl12.Execute(w, players.arrPlayers[currentPIndex])
    if err != nil {
      log.Println(err)
    }
  
  })
  mux.HandleFunc("/leaderboard", leaderboard)
  mux.HandleFunc("/correct", benar)
  mux.HandleFunc("/false", salah)

  log.Fatal(http.ListenAndServe(":9091", mux))
}









