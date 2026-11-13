# ============================================================
# Program  : Luas Permukaan Bola (Kasus G)
# Mata Kuliah : Organisasi dan Arsitektur Komputer (COA)
# Tugas    : T15 Assembly MIPS
# Formula  : L = 4 x phi x r x r  (phi = 3)
# Catatan  : r*r dihitung dengan penjumlahan berulang (loop)
#            4*phi (=12) dikali r*r menggunakan instruksi mul
# ============================================================

        .data
# ---- String identitas kelompok ----
str_header: .asciiz "Kelompok G\nHaydar Errangga Ridwan  103012500117\nMuhammad Raehan Almadani 103012500289\nMuhammad Raihan Marhaban 103012500366\nMuhammad Andra Putra Wicaksana 103012500242"
str_exit:   .asciiz "Untuk keluar dari program, masukkan angka 7007\n\n"

# ---- Prompt input ----
str_r1:     .asciiz "Masukkan Jari-jari Bola 1 : "
str_r2:     .asciiz "Masukkan Jari-jari Bola 2 : "
str_pos:    .asciiz "Bilangan harus positif (> 0). Silakan masukkan ulang.\n"

# ---- Output luas ----
str_luas1:  .asciiz "Luas Permukaan Bola 1 = "
str_luas2:  .asciiz "Luas Permukaan Bola 2 = "
str_nl:     .asciiz "\n"

# ---- Kategori ----
str_kecil:  .asciiz "Kategori : BOLA 1 LEBIH KECIL\n\n"
str_sama:   .asciiz "Kategori : KEDUA BOLA SAMA BESAR\n\n"
str_besar:  .asciiz "Kategori : BOLA 1 LEBIH BESAR\n\n"

        .text
        .globl main

# ============================================================
# PROGRAM UTAMA
# ============================================================
main:
        # Tampilkan header identitas & petunjuk keluar
        li   $v0, 4
        la   $a0, str_header
        syscall
        li   $v0, 4
        la   $a0, str_exit
        syscall

loop_utama:
        # ---- Panggil MASUKAN untuk Jari-jari 1 ----
        la   $a0, str_r1          # prompt jari-jari 1
        jal  MASUKAN              # hasil valid di $v0
        move $s0, $v0             # s0 = r1

        # Cek kondisi keluar: r1 == 7007
        li   $t0, 7007
        beq  $s0, $t0, selesai

        # ---- Panggil MASUKAN untuk Jari-jari 2 ----
        la   $a0, str_r2          # prompt jari-jari 2
        jal  MASUKAN              # hasil valid di $v0
        move $s1, $v0             # s1 = r2

        # ---- Panggil HITUNG ----
        move $a0, $s0             # a0 = r1
        move $a1, $s1             # a1 = r2
        jal  HITUNG               # s2 = L1, s3 = L2, s4 = kategori (0/1/2)

        # ---- Panggil KELUARAN ----
        jal  KELUARAN

        j    loop_utama           # ulang untuk input berikutnya

selesai:
        li   $v0, 10
        syscall                   # exit

# ============================================================
# PROSEDUR MASUKAN
# Input  : $a0 = alamat string prompt
# Output : $v0 = bilangan valid (> 0)
# Menggunakan $t0, $t1 (caller-saved, aman)
# ============================================================
MASUKAN:
        addi $sp, $sp, -8
        sw   $ra, 4($sp)          # simpan return address
        sw   $a0, 0($sp)          # simpan alamat prompt agar tidak hilang saat loop

input_ulang:
        lw   $a0, 0($sp)          # restore prompt
        li   $v0, 4               # cetak prompt
        syscall

        li   $v0, 5               # baca integer dari user
        syscall
        move $t1, $v0             # t1 = nilai yang dimasukkan

        # Validasi: harus > 0
        bgt  $t1, $zero, input_ok # jika positif, lanjut
        li   $v0, 4
        la   $a0, str_pos         # cetak pesan error
        syscall
        j    input_ulang          # minta input ulang

input_ok:
        move $v0, $t1             # kembalikan nilai valid di $v0
        lw   $ra, 4($sp)
        addi $sp, $sp, 8
        jr   $ra

# ============================================================
# PROSEDUR HITUNG
# Input  : $a0 = r1, $a1 = r2
# Output : $s2 = L1, $s3 = L2, $s4 = kategori (0=kecil,1=sama,2=besar)
# Formula: L = 4 x 3 x r x r = 12 x r x r
# r*r  : penjumlahan berulang (loop)
# 12*x : instruksi mul
# ============================================================
HITUNG:
        addi $sp, $sp, -4
        sw   $ra, 0($sp)

        # ---------- Hitung r1 * r1 dengan loop ----------
        move $t0, $a0             # t0 = r1 (counter)
        move $t1, $zero           # t1 = akumulator hasil r1*r1
        move $t2, $a0             # t2 = r1 (nilai yang ditambahkan)
loop_r1sq:
        beq  $t0, $zero, done_r1sq
        add  $t1, $t1, $t2        # akumulator += r1
        addi $t0, $t0, -1         # counter--
        j    loop_r1sq
done_r1sq:
        # t1 = r1^2
        li   $t3, 12
        mul  $s2, $t1, $t3        # L1 = 12 * r1^2

        # ---------- Hitung r2 * r2 dengan loop ----------
        move $t0, $a1             # t0 = r2 (counter)
        move $t1, $zero           # t1 = akumulator hasil r2*r2
        move $t2, $a1             # t2 = r2 (nilai yang ditambahkan)
loop_r2sq:
        beq  $t0, $zero, done_r2sq
        add  $t1, $t1, $t2        # akumulator += r2
        addi $t0, $t0, -1
        j    loop_r2sq
done_r2sq:
        # t1 = r2^2
        li   $t3, 12
        mul  $s3, $t1, $t3        # L2 = 12 * r2^2

        # ---------- Tentukan kategori ----------
        # s4: 0 = BOLA1 LEBIH KECIL, 1 = SAMA, 2 = BOLA1 LEBIH BESAR
        li   $s4, 1               # asumsi sama
        blt  $s2, $s3, kat_kecil
        bgt  $s2, $s3, kat_besar
        j    done_hitung
kat_kecil:
        li   $s4, 0
        j    done_hitung
kat_besar:
        li   $s4, 2
done_hitung:
        lw   $ra, 0($sp)
        addi $sp, $sp, 4
        jr   $ra

# ============================================================
# PROSEDUR KELUARAN
# Input  : $s2 = L1, $s3 = L2, $s4 = kategori
# Output : cetak ke konsol
# ============================================================
KELUARAN:
        addi $sp, $sp, -4
        sw   $ra, 0($sp)

        # Cetak Luas Bola 1
        li   $v0, 4
        la   $a0, str_luas1
        syscall
        li   $v0, 1
        move $a0, $s2
        syscall
        li   $v0, 4
        la   $a0, str_nl
        syscall

        # Cetak Luas Bola 2
        li   $v0, 4
        la   $a0, str_luas2
        syscall
        li   $v0, 1
        move $a0, $s3
        syscall
        li   $v0, 4
        la   $a0, str_nl
        syscall

        # Cetak kategori berdasarkan s4
        li   $t0, 0
        beq  $s4, $t0, cetak_kecil
        li   $t0, 1
        beq  $s4, $t0, cetak_sama
        # else: besar
        li   $v0, 4
        la   $a0, str_besar
        syscall
        j    done_keluaran
cetak_kecil:
        li   $v0, 4
        la   $a0, str_kecil
        syscall
        j    done_keluaran
cetak_sama:
        li   $v0, 4
        la   $a0, str_sama
        syscall
done_keluaran:
        lw   $ra, 0($sp)
        addi $sp, $sp, 4
        jr   $ra