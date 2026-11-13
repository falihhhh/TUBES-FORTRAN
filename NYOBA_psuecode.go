procedure hitungAlas (in alas, tinggi : intenger, in/out luas : real) 
algoritma 
    luas == alas * tinggi / 2
endprocedure

program luasSegitiga 
kamus
    n, i, a, t : integer
algorimta
    input(n)
    for i = 1 to n do 
        input(a, t)
        hitungLuas(a, t, hasil)
        output(hasil)
    endfor
endprogram

______________________________________________________________________________________________________________________________________________

procedure Mencari_nilaiE (in n1, n2 : integer, in/out min, max : integer)
algoritma
    if n1 > n2 then
        min < n1
        max < n2 
    else 
        min < n2 
        max < n1
    endif
endprocedure

program nilaiEKSTRIM 
kamus 
    a, b, c, d : integer
    min1, min2, max1, max2, min, max: integer
algoritma
    input(a,b,c,d)
    mencari_nilaiE = (a,b,min1, max1)
    mencari_nilaiE = (c,d,min2, max2)
    mencari_nilaiE = (max1, max2, max)
    mencari_nilaiE = (min1, min2, min)
    output(max, min)
endprogram

_______________________________________________________________________________________________________________________________________________

procedure demonasi (in uang : integer, in/out k10, k2,k1 : integer)
algoritma 
    in/out k10 / 10000
    uang = uang & 10000

    in/out k2 / 2000
    uang = uang & 2000

    in/out k1/ 1000
    uang = uang & 1000
endprocedure

program : demonasi_DUIT_ANJING
kamus 
    uang, k10, k2, k1 : integer
algoritma 
    input(uang, k10, k2, k1)
    output (k10)
    output (k2)
    output (k1)
endprogram

_______________________________________________________________________________________________________________________________________________

procedure PRENclas (in/out teman, in m : integer )
kamus 
    j : integer
algoritma
    for j < = 0 to m-1 do 
        input(a[j])
    endfor
endprocedure

program : TEMAN
kamus 
    type teman : array [0..100] of string 
    t : teman; n : integer 
algoritma 
    input(n)
    tulis(t, n)
    cetak(t, n)
endprogram

_______________________________________________________________________________________________________________________________________________

program : JARAK_titik
kamus
    type titik : <
        x : real -> (float64)
        y : real -> (float64)
    >
    dataTITIK : array[1..2] of titik 
    hasil : real

function AKAR (x real) real 

function hitung_JARAK (t : array[1..2] of titik) real 
algoritma 
    return akar ((t[1].x - t[2].x)^2 + (t[1].y - t[2].y)^2)
algoritma
    input(dataTITIK[1].x, dataTITIK[1].y)
    input(dataTITIK[2].x, dataTITIK[2].y)

    hasil = hitung_JARAK (dataTITIK)
    output (hasil)
end program

_______________________________________________________________________________________________________________________________________________

program : PERSEGI_panjang
kamus  
    type geometri : <
        luas, keliling : integer
    >
    type PERSEGI_panjang
        panjang, lebar : string
        warna : string
        properti : geometri 
    >
    dataPersegi : array[1..1] of persegiPanjang

procedure isiData(in/out a : array [1..1] of PERSEGI_panjang)
algoritma 
    input (A[1].panjang, A[1].lebar, A[1].warna)
endprocedure

proceduere hitung (in/out a : array [1..1] of PERSEGI_panjang)
algoritma
    a[1].properti.luas = a[1].panjang*a[1].lebar
    a[1].properti.keliling = 2*(a[1].panjang + a[1]labar)
endproceduere

algoritma
    isiData (dataPersegi)
    hitung (dataPersegi)
    output (dataPersegi[1].properti.luas, dataPersegi[1].properti.keliling)
endprogram

_______________________________________________________________________________________________________________________________________________

program matchBOLA
kamus 
    type GOL : array[1..100] of integer
    tim1, tim2 : GOL
    byk1, byk2 : integer
    rata1, rata2 : real

proceduere inputData (in/out t : GOL, out n : integer)
kamus
    nilai : integer
algoritma
    n <= 0
    input (nilai)

    while nilai >= 0 do
        n <= n+1
        t[n] <= nilai
        input (nilai)
    endwhile
endproceduere

function RATA (t : gol, n : integer) real
kamus
    i, total : integer
algoritma
    total = 0
    for i = 1 to n do 
        total = total + t[i]
    endfor
    
    if n = 0 then
        return 0 
    else 
        return total / new
    endif
endfunction

algoritma
    inputData (tim1, byk1)
    inutData (tim2, byk2)

    rata1 = RATA (tim1, byk1)
    rata1 = RATA (tim2, byk2)

    output(rata1, rata2)
endprogram

_______________________________________________________________________________________________________________________________________________

//LATIHAN UTS

//SOAL UTS NO 1 

//BAGIAN : A

function biayaParcel (berat, biayaPerGRAM, diskon, pajak : real) real {
kamus
    biayaAwal, biayaAFTERdiskon, TOTALbiaya : real
algoritma 
    biayaAwal = berat * 1000 * biayaPerGRAM
    biayaAFTERdikson = biayaAwal - (biayaAwal * dikson / 100)
    TOTALbiaya = biayaAFTERdiskon + (biayaAFTERdiskon * pajak / 100)
    return TOTALbiaya
} //(endfunction)

//BAGIIAN : B

procedure hitungWAKTU (jarak, kecepatan : real in/out hari : integer, in/out jam : real) 
kamus
    TOTALjam : real
algoritma
    TOTALjam = jarak / kecepatan
    hari = 0
    while TOTALjam >= 24 do
        hari = hari + 1
        TOTALjam = TOTALjam - 24
    endwhile
    jam = TOTALjam
endprocedure

//SOAL UTS NO 2

function GDP (k, i, g, x, m : integer) integer
kamus 
    hasil : integer
algoritma
    hasil = k + i + g + x + w
    return hasil
endfunction

procedure hitungRerata (in N : integer, in/out rata2 : real)
kamus
        j, k, i, g, x, m, nilaiGDP : integer
    totalGDP : real

algoritma
    totalGDP = 0
    for j = 1 to N do
        input (k, i, g, x, m)
            nilaiGDP = gdp (k, i, g, x, m)
            output (nilaiGDP, " trillion")
            totalGDP = totalGDP + nilaiGDP
        endfor
    rata2 = totalGDP / N 
endprocedure

Program : GROSS PRODUCT
kamus 
    n : integer
    avarage : real
algoritma
    input(n)
    hitungAVARAGE (n, rata2)
    output ("Avarage GDP: ", rata2, " trillion")
endprogram

//SOAL UTS NO 3

function DIGIT (n : integer)integer
kamus  
    LASTdigit : integer
algoritma
    LASTdigit = n mod 10

    if n < 10 then
        if LASTdigit mod 2 = 1 then 
            return 1
        else
            return 0
        endif
    else
        if LASTdigit mod 2 = 1 then
            return 1 + DIGIT (n div 10)
        else
            return DIGIT (n div 10)
        endif 
    endif
endfunction

Program : NUMBER
kamus 
    n, hasil : integer
    input (n)
    hasil = DIGIT (n)
    output (hasil)
endprogram

//SOAL TRACING

program ABC
algoritma
    output (BCD(4))
endprogram
function BCD(p : integer) à integer
kamus
    q, k : integer
algoritma
    for k = 1 to p do
        DEF(k, q)
    endfor
    return q
endfunction

procedure DEF(in x : integer, in/outy : integer)
kamus
    j : integer
algoritma
    for j = 1 to x do y = j * x
        output(j + y)
    endfor
endprocedure
