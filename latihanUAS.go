//LATIHAN SOAL ALGORITMA DAN PEMROGRAMAN 2

//Materi: Searching (Sequential & Binary) dan Sorting (Selection & Insertion)

//SOAL 1: Pencarian Data (Searching)

//Diberikan sebuah tipe bentukan Mahasiswa yang menyimpan informasi NIM (string) dan Nama (string). Data seluruh mahasiswa disimpan dalam sebuah array tabMhs berkapasitas NMAX.

//Pertanyaan:
//Buatlah dua buah fungsi untuk mencari indeks array di mana mahasiswa dengan NIM tertentu (sebut saja X) berada! Jika tidak ditemukan, kembalikan nilai -1.
//a. Buatlah fungsi cariNIM_Seq menggunakan algoritma Sequential Search (asumsi data pada array acak).
//b. Buatlah fungsi cariNIM_Bin menggunakan algoritma Binary Search (asumsi data pada array sudah terurut Ascending berdasarkan NIM).

//JAWABAN SEQURNTIAL SEARCH 
func cariNIM (T  : TABmhs, N : integer, X : string)integer
kamus
	i : integer
algoritma 
	for i = 0 to N-1 do
		if T[i].NIM == X then 
			return i
		endif
	endfor

	return -1
endfunc


//JAWABAN BINARY SEARCH
func cariNIM (T : TABmhs, N : integer, S : string)integer
kamus 
	kiri, kanan, tengah : integer
algoritma
	kiri = 0 
	kanan = N-1

	while kiri <= kanan do
		tengah = (kiri + kanan) div 2

		if T[tengah].NIM == X then 
			return tengah
		else if X < T[tengah].NIM then
			kanan = tengah - 1
		else 
			kiri = tengah + 1
		endif
	endwhile

	return -1
endfunc

//SOAL 2: Pengurutan Data (Sorting)

//Diberikan sebuah tipe bentukan Peserta yang menyimpan informasi Nama (string) dan Skor (integer). Data peserta lomba disimpan dalam sebuah array tabPeserta sebanyak N data.

//Pertanyaan:
//Buatlah dua buah prosedur untuk mengurutkan array tersebut berdasarkan Skor secara Descending (mengecil, dari skor tertinggi ke skor terendah)!
//a. Buatlah prosedur urutSkor_Selection menggunakan algoritma Selection Sort.
//b. Buatlah prosedur urutSkor_Insertion menggunakan algoritma Insertion Sort.

//JAWABAN SOAL 2

//a) Jawaban Selection Sort (Descending)
procedure urutSKOR (in/out T : TABpeserta,in N : integer)
kamus 
	pass, j, MAXidx : integer
	temp : peserta
algoritma 
	for pass = 0 to N-2 ddo
		MAXidx = pass

		for j = pass + 1 to N-1 do
			if T[j].skor > T[MAXidx].skor then 
				MAXidx = j
			endif
		endfor

		temp = T[pass]
		pass = T[MAXidx]
		T[MAXidx] = temp
	endfor
endprocedure

//b) Jawaban Insertion Sort (Descending):
proceduere urutSKOR (in/out T : tabPROCEDUREpeserta, in N : integer)
kamus 
	pass, j : integer
	temp : peserta
algoritma
	for pass = 1 to N-1 do
		temp = T[pass]
		j = pass -1

		while (j >= 0) and (T[j].skor < temp.Skor) do
			T[t+j] = T[j]
			j < j-1
		endwhile

		T[j+1] =  temp
	endfor
endprocedure


==================================================================================================================================================================================================================


//LATIHAN SOAL ALGORITMA DAN PEMROGRAMAN 2 (VERSI 2)
//Materi: Searching (Sequential & Binary) dan Sorting (Selection & Insertion)

//SOAL 1: Pencarian Data (Searching)
//Diberikan sebuah tipe bentukan Karyawan yang menyimpan informasi IDPegawai (string) dan Gaji (integer). Data seluruh karyawan di sebuah perusahaan disimpan dalam array tabKaryawan yang berisi N data.

//Pertanyaan:
//Buatlah dua buah fungsi untuk mencari indeks array di mana karyawan dengan IDPegawai tertentu (sebut saja X) berada! Jika tidak ditemukan, kembalikan nilai -1.
//a. Buatlah fungsi cariID_Seq menggunakan algoritma Sequential Search (asumsi data pada array acak).
//b. Buatlah fungsi cariID_Bin menggunakan algoritma Binary Search (asumsi data pada array sudah terurut Descending berdasarkan IDPegawai). (Hati-hati: Asumsi data terurut Descending)

//JAWABAN SOAL 1

//a) Jawaban Sequential Search:
func cariID (T  : TabKaryawan, N : integer, X : integer)integer
kamus
	i : integer
algoritma 
	for i = 0 to N-1 do
	 	if T[i].IDpengawal == x then
			return i
		endif
	endfor 

	return i
endfunc

//b) Jawaban Binary Search (Data terurut Descending):
func cariID (T : tabKaryawan, N : integer, X : string) integer
kamus 
	kiri, kanan, tengah : integer
algoritma
	kiri = 0
	kanan = 0

	while kiri <= kanan do 
		tengaj = (kiri + kanan) div 2

		if T[tengah].IDpengawal == X then
			return tengah
		else if X > T[tengah]/IDpengawal then 
			kanan = tengah -1
		else 
			kiri = tengah
		endif
	endwhile

	return i 
endfunc

//SOAL 2: Pengurutan Data (Sorting)

//Diberikan sebuah tipe bentukan Barang yang menyimpan informasi KodeBarang (string) dan Stok (integer). Data barang di gudang disimpan dalam array tabBarang sebanyak N data.
//Pertanyaan:
//Buatlah dua buah prosedur untuk mengurutkan array tersebut berdasarkan Stok secara Ascending (membesar, dari stok terdikit ke stok terbanyak)!
//a. Buatlah prosedur urutStok_Selection menggunakan algoritma Selection Sort.
//b. Buatlah prosedur urutStok_Insertion menggunakan algoritma Insertion Sort.

//JAWABAN SOAL 2

//a) Jawaban Selection Sort (Ascending):
proceduere urutSTOK (in/out T: tabBARANG, in N : inetgerd)
kamus
	pass, j, MINdix : integer
	temp : barang
algoritma 
	for psas = 0 to N-1 do
		MINidx = pass

		for j = pas + 1 to N-1 do
			if T[j].stok < T[MINidx].stok then
				MINidx = j
			endif
		endfor

		temp = [pass]
		pass = T[MINidx]
		T[MINidx] = temp
	endfor
endprocedure

//b) Jawaban Insertion Sort (Ascending):
procedure urutSTOK (in/out T : tabBARANG, in n : integer)
kamus 
	pass, j : integer
	temp :  Barang 
algoritma
	for pass = 1 to N-1 do
		temp = T[pass]
		j = pass - 1 

		while (j >= 0) and (T[j].stok > temp.stok) do
			T[t+1]= temp
			j = j = 1
		endwhile

		T[j+1] = temp
	endfor
endprocedure



//LATIHAN UTS DAN QUIZ C++/CPP

LATIHAN SOAL UTS 2025
1A. Pembuatan ADT SLL dan DLL (Integer)
Singly Linked List (SLL):

type Infotype : integer
type Address : pointer to ElmList
type ElmList < 
	info : Infotype 
	next : Address 
	>
type SinglyList < 
	first : Address
	>

Doubly Linked List (DLL):
type Infotype : integer
type Address : pointer to ElmList
type ElmList < 
	info : Infotype
	next : Address
	prev : Address 
	 >
type DoublyList < 
	first : Address
	last : Address
	>

1B. PROCEDURE InsertlastSLL
algorithm
    if L.first = NIL then
		L.first = p
	else 
		Q <- L.first
		while Q->next != NIL do
			Q <- Q->next
			Q->next <- P
		end while
	endif	

1C. PROCEDURE insertFirst DLL
algorithm
    if L.first = NIL then
        L.first <- P
        L.last <- P
    else
        P->next <- L.first
        L.first->prev <- P
        L.first <- P
    endif
endprocedure

1D. Analisis PROCEDURE PopStack [Array]
	procedure pop( in/out S: Stack, in/out P: Infotype )
	{I.S. terdefinisi sebuah Stack S yang mungkin kosong dan P disiapkan untuk menampung hasil pop}
	{F.S. elemen yang dikeluarkan akan disimpan di P}
	kamus
	Procedure isEmpty( Stack ) -> Boolean
	algoritma
		if not isEmpty( S ) then
			P = S.info[ S.top ]
			S.top = S.top + 1 
		endif
	endprogram
____________________________________________________________________________________________________
Analisis :
	"S.top = S.top + 1" = salah, karna fungsi untuk mengahapus/ambil elemen stack
	seharusnya jumlah elemen berkurang bukan bertambah jadi kode yg benar adalah 
	"S.top = S.top - 1 "

	"P = S.info[ S.top ] (atau P = S[S.top]" = salah, karna seharunys simbol "="
	ditulis "<-" tapi ini bisa berubah ubah sesuai ketentuan dosen

	setelah "endif" ada "endprogram" di kode, ini salah karna harusnya "endproceduere"
	aku di atas ga bikin main program dan seharunya "endprocedure" bukan "endprogram"

kode yg benar : 
	procedure pop(in/out S: Stack, in/out P: Infotype)
	{I.S. terdefinisi sebuah Stack S yang mungkin kosong dan P disiapkan untuk menampung hasil pop}
	{F.S. elemen yang dikeluarkan akan disimpan di P dan top berkurang 1}
	kamus
    Procedure isEmpty(Stack) -> Boolean
	algoritma
    	if not isEmpty(S) then
        	P <- S.info[S.top]
        	S.top <- S.top - 1
    	endif
	endprocedure

1E. 1e. Analisis dan Koreksi Procedure enqueue Queue [DLL]
	procedure enqueue( in/out Q: Queue, in P: Address )
	{I.S. terdefinisi sebuah Queue Q yang mungkin kosong dan sebuah pointer P yang menunjuk ke sebuah elemen baru}
	{F.S. elemen yang ditunjuk P akan masuk ke dalam Q}
	kamus
	Procedure isEmpty( Queue ) -> Boolean
	algoritma
		if isEmpty( Q ) then
			Q.head = Q.tail = P
		else
			P->prev = Q.tail
			Q.tail->prev = P 
			Q.tail = P
		endif
	endprogram
____________________________________________________________________________________________________
Analisis :
	"Q.tail -> prev = P" =  sa;ah, karna nambah elmen baru dibelekang ekor lama "Q.tail"
	hrs dihubungkan penunjuk arah maju "next" dari yg lama ke elemen baru "p"
	mengubah "prev" dari ekor lama akan merusak Linked list nya, yg benar seharusnya
	seperti ini "Q.tail->next <- P"

	"P -> prev = Q.tail dan Q.tail = P" = salah semua karna, operasi nilai menggunakan
	"=", memasukan nilai wajib menggunakan "<-" tapi ini sesuai ketentuan dosen ygy

	"Q.head = Q.tail = P" = salah karna, ganda dalam satu baris tidak sah dalam
	aturan psuedcode, jadi yg benar harusnya "Q.head <- P dan Q.tail <- P"

	setelah "endif" ada "endprogram" di kode, ini salah karna harusnya "endproceduere"
	aku di atas ga bikin main program dan seharunya "endprocedure" bukan "endprogram"

Kode yang banar :
	procedure enqueue(in/out Q: Queue, in P: Address)
	{I.S. terdefinisi sebuah Queue Q yang mungkin kosong dan sebuah pointer P yang menunjuk ke sebuah elemen baru}
	{F.S. elemen yang ditunjuk P akan masuk ke dalam Q}
	kamus
    	procedure isEmpty(Queue) -> Boolean
	algoritma
    	if isEmpty(Q) then
        	Q.head <- P
        	Q.tail <- P
    	else
        	P->prev <- Q.tail
        	Q.tail->next <- P
        	Q.tail <- P
    	endif
	endprocedure


2. SLL
Type Infotype : Lagu
Type Address : pointer to ElmList
Type Lagu : < judul: string,
             penyanyi: string,
             album: string,
             durasi: integer 
			>
procedure separateBySinger(in/out L: List, in/out L1: List, in/out L2: List, in/out L3: List)
{ I.S. Terdefinisi list L berisi lagu-lagu dari 3 penyanyi, L1, L2, L3 adalah list kosong }
{ F.S. Elemen dari L dipindahkan ke L1, L2, dan L3 sesuai penyanyinya. L menjadi kosong }
kamus
    P: Address
    nama1, nama2, nama3: string
algoritma
    nama1 <- ""
    nama2 <- ""
    nama3 <- ""
 
    while not isEmpty(L) do
        deleteFirst(L, P)
        
        if nama1 = "" or nama1 = P->info.penyanyi then
            nama1 <- P->info.penyanyi
            insertLast(L1, P)
        else if nama2 = "" or nama2 = P->info.penyanyi then
            nama2 <- P->info.penyanyi
            insertLast(L2, P)
        else
            nama3 <- P->info.penyanyi
            insertLast(L3, P)
        endif
    endwhile
endprocedure

3. DLL

4. STACK 
procedure popByValue(in/out S: Stack, in X: Infotype)
{ I.S. terdefinisi sebuah Stack S yang mungkin kosong dan nilai target X }
{ F.S. jika X ditemukan, X dihapus dari S. Elemen lain tetap dalam urutan semula. Jika tidak, cetak pesan }
kamus
    tempS : Stack     
    val : Infotype    
    ketemu : Boolean  
algoritma
    createStack(tempS)
    ketemu <- false
    
    while (not isEmpty(S)) and (not ketemu) do
        pop(S, val)
        if val = X then
            ketemu <- true 
        else
            push(tempS, val) 
        endif
    endwhile
    
    while not isEmpty(tempS) do
        pop(tempS, val)
        push(S, val)
    endwhile

    if not ketemu then
        output("Elemen tidak ditemukan di dalam Stack")
    endif
endprocedure