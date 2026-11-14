@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@400;600;700;800&family=Plus+Jakarta+Sans:wght@400;500;600;700&display=swap');

:root {
    --dark-bg: #1b1d1f;
    --dark-card: #24272a;
    --light-bg: #ffffff;
    --light-card: #f8f9fa;
    --accent: #36444f;
    --accent-hover: #253038;
    --accent-light: #5c7285;
    --text-white: #ffffff;
    --text-gray-light: #a1a1aa;
    --text-black: #18181b;
    --text-gray-dark: #52525b;
    --border-dark: #32363a;
    --border-light: #e4e4e7;
}

* {
    box-sizing: border-box;
    margin: 0;
    padding: 0;
}

html {
    scroll-behavior: smooth;
}

body {
    font-family: 'Plus Jakarta Sans', sans-serif;
    background-color: var(--dark-bg);
    color: var(--text-white);
    line-height: 1.7;
    overflow-x: hidden;
}

/* Animasi CSS Murni */
@keyframes fadeUp {
    from {
        opacity: 0;
        transform: translateY(24px);
    }
    to {
        opacity: 1;
        transform: translateY(0);
    }
}

.animate-in {
    animation: fadeUp 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
}

/* Struktur Band Selang-Seling (Dark & Light) */
.band-dark {
    background-color: var(--dark-bg);
    color: var(--text-white);
    padding: 75px 20px;
}

.band-light {
    background-color: var(--light-bg);
    color: var(--text-black);
    padding: 75px 20px;
}

.container {
    max-width: 1000px;
    margin: 0 auto;
}

/* 1. HERO SECTION (DARK) */
.hero {
    text-align: center;
    padding: 90px 20px 85px 20px;
    border-bottom: 1px solid var(--border-dark);
}

.hero-kicker {
    font-family: 'Montserrat', sans-serif;
    font-size: 0.78rem;
    font-weight: 700;
    letter-spacing: 3px;
    text-transform: uppercase;
    color: var(--text-gray-light);
    margin-bottom: 14px;
}

.hero h1 {
    font-family: 'Montserrat', sans-serif;
    font-size: 2.1rem;
    font-weight: 800;
    text-transform: uppercase;
    letter-spacing: 1px;
    line-height: 1.3;
    max-width: 850px;
    margin: 0 auto 32px auto;
}

/* Tombol Ganda di Hero ("Lihat Hasil" & "Coba Demo") */
.hero-actions {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 14px;
    flex-wrap: wrap;
}

.btn-red {
    display: inline-block;
    background-color: var(--accent);
    color: #ffffff;
    border: 1px solid var(--accent-light);
    font-family: 'Montserrat', sans-serif;
    font-size: 0.75rem;
    font-weight: 700;
    letter-spacing: 2px;
    text-transform: uppercase;
    text-decoration: none;
    padding: 13px 28px;
    border-radius: 3px;
    cursor: pointer;
    transition: transform 0.25s ease, background-color 0.25s ease, box-shadow 0.25s ease, border-color 0.25s ease;
}

.btn-red:hover {
    background-color: var(--accent-hover);
    border-color: #7a93a8;
    transform: translateY(-3px);
    box-shadow: 0 8px 20px rgba(54, 68, 79, 0.55);
}

.btn-outline {
    display: inline-block;
    background-color: transparent;
    color: #ffffff;
    border: 1px solid rgba(255, 255, 255, 0.35);
    font-family: 'Montserrat', sans-serif;
    font-size: 0.75rem;
    font-weight: 700;
    letter-spacing: 2px;
    text-transform: uppercase;
    text-decoration: none;
    padding: 13px 28px;
    border-radius: 3px;
    cursor: pointer;
    transition: transform 0.25s ease, background-color 0.25s ease, border-color 0.25s ease, color 0.25s ease;
}

.btn-outline:hover {
    background-color: #ffffff;
    color: #18181b;
    border-color: #ffffff;
    transform: translateY(-3px);
}

/* Heading Ala Editorial */
.section-header {
    text-align: center;
    margin-bottom: 45px;
}

.section-kicker {
    font-family: 'Montserrat', sans-serif;
    font-size: 0.72rem;
    font-weight: 700;
    letter-spacing: 3px;
    text-transform: uppercase;
    color: var(--text-gray-dark);
    margin-bottom: 6px;
}

.band-dark .section-kicker {
    color: var(--text-gray-light);
}

.section-title-wrap {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 20px;
}

.section-title-wrap::before,
.section-title-wrap::after {
    content: "";
    height: 1px;
    width: 70px;
    background-color: currentColor;
    opacity: 0.25;
}

.section-header h2 {
    font-family: 'Montserrat', sans-serif;
    font-size: 1.65rem;
    font-weight: 800;
    text-transform: uppercase;
    letter-spacing: 2px;
}

.section-desc {
    max-width: 720px;
    margin: 14px auto 0 auto;
    font-size: 0.92rem;
    color: var(--text-gray-dark);
}

.band-dark .section-desc {
    color: var(--text-gray-light);
}

/* 3 Kolom Highlight Ringkasan */
.grid-3 {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(270px, 1fr));
    gap: 28px;
    margin-bottom: 50px;
}

.editorial-card {
    padding: 24px;
    border: 1px solid var(--border-light);
    background: var(--light-card);
    transition: transform 0.3s ease, border-color 0.3s ease;
}

.band-dark .editorial-card {
    background: var(--dark-card);
    border-color: var(--border-dark);
}

.editorial-card:hover {
    transform: translateY(-5px);
    border-color: var(--accent);
}

.band-dark .editorial-card:hover {
    border-color: var(--accent-light);
}

.card-kicker {
    font-family: 'Montserrat', sans-serif;
    font-size: 0.7rem;
    font-weight: 700;
    letter-spacing: 2px;
    text-transform: uppercase;
    opacity: 0.65;
}

.editorial-card h3 {
    font-family: 'Montserrat', sans-serif;
    font-size: 1.15rem;
    font-weight: 800;
    text-transform: uppercase;
    letter-spacing: 1px;
    margin: 4px 0 12px 0;
}

.red-bar {
    width: 32px;
    height: 3px;
    background-color: var(--accent);
    margin-bottom: 14px;
    transition: width 0.3s ease;
}

.band-dark .red-bar {
    background-color: var(--accent-light);
}

.editorial-card:hover .red-bar {
    width: 60px;
}

.editorial-card p {
    font-size: 0.88rem;
    line-height: 1.6;
    opacity: 0.85;
}

/* Tabel Monokrom Minimalis */
.sub-heading {
    font-family: 'Montserrat', sans-serif;
    font-size: 1.05rem;
    font-weight: 700;
    text-transform: uppercase;
    letter-spacing: 1px;
    margin: 36px 0 14px 0;
}

.table-wrapper {
    overflow-x: auto;
    margin-bottom: 28px;
    border: 1px solid var(--border-light);
}

.band-dark .table-wrapper {
    border-color: var(--border-dark);
}

table {
    width: 100%;
    border-collapse: collapse;
    font-size: 0.9rem;
}

th {
    font-family: 'Montserrat', sans-serif;
    font-size: 0.78rem;
    text-transform: uppercase;
    letter-spacing: 1px;
    padding: 14px 16px;
    text-align: center;
    background-color: #18181b;
    color: #ffffff;
}

.band-dark th {
    background-color: #121315;
    border-bottom: 1px solid var(--border-dark);
}

td {
    padding: 13px 16px;
    text-align: center;
    border-bottom: 1px solid var(--border-light);
}

.band-dark td {
    border-bottom: 1px solid var(--border-dark);
}

th:first-child, td:first-child {
    text-align: left;
    font-weight: 600;
}

tr {
    transition: background-color 0.2s ease;
}

.band-light tr:hover td {
    background-color: #f4f4f5;
}

.band-dark tr:hover td {
    background-color: #24272a;
}

.mono-bold {
    font-weight: 800;
    text-decoration: underline;
    text-decoration-color: var(--accent);
    text-underline-offset: 4px;
    text-decoration-thickness: 2px;
}

.band-dark .mono-bold {
    text-decoration-color: var(--accent-light);
}

.mono-muted {
    opacity: 0.65;
}

/* Kotak Grafik */
.chart-box {
    padding: 24px;
    border: 1px solid var(--border-light);
    background: #ffffff;
    margin: 24px 0;
}

.band-dark .chart-box {
    background: var(--dark-card);
    border-color: var(--border-dark);
}

/* Box Analisis Editorial */
.analysis-note {
    border-left: 3px solid var(--accent);
    padding: 14px 20px;
    margin: 20px 0;
    background: rgba(128, 128, 128, 0.06);
    font-size: 0.92rem;
}

.band-dark .analysis-note {
    border-left-color: var(--accent-light);
}

.analysis-note ul {
    padding-left: 18px;
    margin-top: 8px;
}

.analysis-note li {
    margin-bottom: 6px;
}

/* FITUR INTERAKTIF: COBA SENDIRI (MONOCHROME + #36444f) */
.demo-box {
    background: var(--light-card);
    border: 1px solid var(--border-light);
    padding: 32px;
    border-radius: 4px;
}

.demo-group {
    margin-bottom: 22px;
}

.demo-label {
    display: block;
    font-family: 'Montserrat', sans-serif;
    font-size: 0.75rem;
    font-weight: 700;
    text-transform: uppercase;
    letter-spacing: 1px;
    color: var(--text-gray-dark);
    margin-bottom: 8px;
}

.demo-textarea,
.demo-input,
.demo-select {
    width: 100%;
    padding: 12px 14px;
    font-family: 'Plus Jakarta Sans', sans-serif;
    font-size: 0.95rem;
    color: var(--text-black);
    background: #ffffff;
    border: 1px solid #d4d4d8;
    border-radius: 3px;
    transition: border-color 0.2s ease, box-shadow 0.2s ease;
}

.demo-textarea:focus,
.demo-input:focus,
.demo-select:focus {
    outline: none;
    border-color: var(--accent);
    box-shadow: 0 0 0 3px rgba(54, 68, 79, 0.14);
}

.demo-textarea {
    min-height: 76px;
    resize: vertical;
}

.demo-row {
    display: grid;
    grid-template-columns: 1fr 1fr auto;
    gap: 16px;
    align-items: end;
    padding-top: 18px;
    border-top: 1px solid var(--border-light);
}

.btn-dark-solid {
    background-color: #18181b;
    color: #ffffff;
    border: 1px solid #18181b;
    font-family: 'Montserrat', sans-serif;
    font-size: 0.75rem;
    font-weight: 700;
    letter-spacing: 1.5px;
    text-transform: uppercase;
    padding: 13px 24px;
    border-radius: 3px;
    cursor: pointer;
    transition: background-color 0.25s ease, transform 0.2s ease;
}

.btn-dark-solid:hover {
    background-color: #3f3f46;
    transform: translateY(-2px);
}

.demo-status {
    margin: 24px 0 18px 0;
    font-size: 0.9rem;
    font-weight: 600;
    color: var(--text-gray-dark);
}

.demo-results {
    display: flex;
    flex-direction: column;
    gap: 18px;
}

.algo-row {
    display: grid;
    grid-template-columns: 150px 1fr;
    gap: 16px;
    align-items: start;
}

.algo-name {
    font-family: 'Montserrat', sans-serif;
    font-size: 0.85rem;
    font-weight: 800;
    text-transform: uppercase;
    letter-spacing: 0.5px;
    padding-top: 2px;
}

.algo-name.bubble { color: var(--accent); }
.algo-name.selection { color: #18181b; }
.algo-name.insertion { color: #52525b; }

.bar-track {
    width: 100%;
    height: 12px;
    background-color: #e4e4e7;
    border-radius: 999px;
    overflow: hidden;
    margin-bottom: 6px;
}

.bar-fill {
    height: 100%;
    border-radius: 999px;
    transition: width 0.55s cubic-bezier(0.16, 1, 0.3, 1);
}

.bar-fill.bubble { background-color: var(--accent); }
.bar-fill.selection { background-color: #18181b; }
.bar-fill.insertion { background-color: #71717a; }

.algo-meta {
    font-size: 0.85rem;
    color: var(--text-gray-dark);
}

.fastest-badge {
    color: var(--accent);
    font-weight: 800;
}

/* Kesimpulan Grid */
.conclusion-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
    gap: 24px;
}

footer {
    background-color: #121315;
    color: var(--text-gray-light);
    text-align: center;
    padding: 30px 20px;
    font-family: 'Montserrat', sans-serif;
    font-size: 0.75rem;
    letter-spacing: 2px;
    text-transform: uppercase;
    border-top: 1px solid var(--border-dark);
}

@media (max-width: 720px) {
    .hero h1 { font-size: 1.5rem; }
    .section-title-wrap::before, .section-title-wrap::after { width: 25px; }
    .demo-row { grid-template-columns: 1fr; }
    .algo-row { grid-template-columns: 1fr; gap: 6px; }
    .conclusion-grid { grid-template-columns: 1fr; }
}























<!DOCTYPE html>
<html lang="id">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Analisis Algoritma Sorting — Kelompok 26 PELLINORE</title>
    <link rel="stylesheet" href="style.css">
    <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
</head>
<body>

    <!-- BAND 1: HERO SECTION (DARK) -->
    <header class="band-dark hero">
        <div class="container animate-in">
            <div class="hero-kicker">Kelompok 26 — Pellinore • Tugas Besar Fortran 2026</div>
            <h1>Analisis Perbandingan Kecepatan Algoritma Bubble Sort, Selection Sort, dan Insertion Sort pada Berbagai Ukuran dan Kondisi Data</h1>
            <div class="hero-actions">
                <a href="#hasil-pengujian" class="btn-red">Lihat Hasil</a>
                <a href="#coba-demo" class="btn-outline">Coba Demo</a>
            </div>
        </div>
    </header>

    <!-- BAND 2: HASIL DATA ACAK (LIGHT) -->
    <section id="hasil-pengujian" class="band-light">
        <div class="container animate-in">
            <div class="section-header">
                <div class="section-kicker">Empirical Benchmark</div>
                <div class="section-title-wrap">
                    <h2>Hasil Pengujian</h2>
                </div>
                <p class="section-desc">Pengukuran waktu eksekusi empiris dalam milidetik (ms) berdasarkan rata-rata dan simpangan baku (SD) dari 5 kali pengulangan.</p>
            </div>

            <!-- 3 Kartu Sorotan -->
            <div class="grid-3">
                <div class="editorial-card">
                    <div class="card-kicker">Data Acak • n = 5.000</div>
                    <h3>Selection Tercepat</h3>
                    <div class="red-bar"></div>
                    <p>Mencatat waktu 327,49 ms — sekitar 2,5 kali lebih cepat dibanding Bubble Sort (825,30 ms).</p>
                </div>
                <div class="editorial-card">
                    <div class="card-kicker">Best Case • Terurut</div>
                    <h3>Perilaku Linier O(n)</h3>
                    <div class="red-bar"></div>
                    <p>Bubble Sort (dengan penanda) dan Insertion Sort selesai di bawah 0,2 ms, ribuan kali lebih cepat dari data acak.</p>
                </div>
                <div class="editorial-card">
                    <div class="card-kicker">Worst Case • Terbalik</div>
                    <h3>Minim Pertukaran</h3>
                    <div class="red-bar"></div>
                    <p>Selection Sort unggul (51,87 ms) karena jumlah pertukaran maksimal hanya n − 1 kali.</p>
                </div>
            </div>

            <!-- Tabel 2 -->
            <h3 class="sub-heading">A. Pengaruh Ukuran Data pada Data Acak (Tabel 2)</h3>
            <div class="table-wrapper">
                <table>
                    <thead>
                        <tr>
                            <th>Ukuran Data (n)</th>
                            <th>Bubble Sort (ms)</th>
                            <th>Selection Sort (ms)</th>
                            <th>Insertion Sort (ms)</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr><td>100</td><td>0,31 ± 0,10</td><td>0,11 ± 0,00</td><td><span class="mono-bold">0,10 ± 0,01</span></td></tr>
                        <tr><td>500</td><td>6,39 ± 0,11</td><td>3,04 ± 0,05</td><td><span class="mono-bold">2,93 ± 0,20</span></td></tr>
                        <tr><td>1.000</td><td>27,43 ± 0,88</td><td>12,18 ± 0,52</td><td><span class="mono-bold">11,80 ± 0,25</span></td></tr>
                        <tr><td>2.500</td><td>193,64 ± 10,34</td><td>83,29 ± 7,14</td><td><span class="mono-bold">80,66 ± 4,99</span></td></tr>
                        <tr><td>5.000</td><td><span class="mono-muted">825,30 ± 64,58</span></td><td><span class="mono-bold">327,49 ± 40,11</span></td><td>352,91 ± 62,45</td></tr>
                    </tbody>
                </table>
            </div>

            <div class="chart-box">
                <canvas id="randomChart"></canvas>
            </div>

            <div class="analysis-note">
                <strong>Analisis Pertumbuhan Waktu Eksekusi:</strong>
                <ul>
                    <li><strong>Pertumbuhan Kuadratik O(n²):</strong> Ketika <i>n</i> naik 5 kali lipat (1.000 ke 5.000), waktu eksekusi melonjak 27–30 kali lipat (Bubble Sort 30,1x, Selection Sort 26,9x, Insertion Sort 29,9x), mendekati prediksi teoretis 5² = 25.</li>
                    <li><strong>Bubble Sort Paling Lambat:</strong> Pada <i>n</i> = 5.000 membutuhkan 825,30 ms (2,5x lebih lama dari Selection Sort dan 2,3x lebih lama dari Insertion Sort).</li>
                    <li><strong>Selection vs Insertion Sort:</strong> Performa keduanya hampir setara; selisih pada <i>n</i> = 5.000 masih berada dalam rentang simpangan baku.</li>
                </ul>
            </div>
        </div>
    </section>

    <!-- BAND 3: KONDISI DATA & PERBANDINGAN (DARK) -->
    <section class="band-dark">
        <div class="container">
            <div class="section-header">
                <div class="section-kicker">Data Sensitivity & Operations</div>
                <div class="section-title-wrap">
                    <h2>Kondisi & Operasi Data</h2>
                </div>
                <p class="section-desc">Evaluasi sensitivitas algoritma terhadap urutan awal masukan (n = 2.000) serta jumlah perbandingan elemen (n = 1.000).</p>
            </div>

            <!-- Tabel 3 -->
            <h3 class="sub-heading">B. Pengaruh Kondisi Data pada n = 2.000 (Tabel 3)</h3>
            <div class="table-wrapper">
                <table>
                    <thead>
                        <tr>
                            <th>Kondisi Data</th>
                            <th>Bubble Sort (ms)</th>
                            <th>Selection Sort (ms)</th>
                            <th>Insertion Sort (ms)</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>Terurut (Best Case)</td>
                            <td><span class="mono-bold">0,07 ± 0,01</span></td>
                            <td><span class="mono-muted">48,38 ± 1,63</span></td>
                            <td><span class="mono-bold">0,12 ± 0,01</span></td>
                        </tr>
                        <tr>
                            <td>Acak (Average Case)</td>
                            <td><span class="mono-muted">113,45 ± 3,98</span></td>
                            <td><span class="mono-bold">46,81 ± 0,78</span></td>
                            <td>58,49 ± 14,17</td>
                        </tr>
                        <tr>
                            <td>Terbalik (Worst Case)</td>
                            <td><span class="mono-muted">157,24 ± 27,85</span></td>
                            <td><span class="mono-bold">51,87 ± 2,83</span></td>
                            <td>89,16 ± 1,73</td>
                        </tr>
                    </tbody>
                </table>
            </div>

            <div class="chart-box">
                <canvas id="casesChart"></canvas>
            </div>

            <div class="analysis-note">
                <strong>Analisis Pengaruh Kondisi Data:</strong>
                <ul>
                    <li><strong>Data Terurut:</strong> Bubble Sort (dengan penanda pertukaran) dan Insertion Sort selesai dalam &lt;0,2 ms karena berperilaku O(n) — lebih cepat &gt;1.600x dan ~500x dibanding data acak. Selection Sort tidak terpengaruh (48,38 ms).</li>
                    <li><strong>Data Terbalik:</strong> Waktu Bubble Sort naik 1,4x dan Insertion Sort naik 1,5x dari data acak. Selection Sort menjadi yang tercepat (51,87 ms) karena jumlah pertukarannya paling sedikit.</li>
                </ul>
            </div>

            <!-- Tabel 4 -->
            <h3 class="sub-heading">C. Jumlah Perbandingan Antar-Elemen pada n = 1.000 Data Acak (Tabel 4)</h3>
            <div class="table-wrapper">
                <table>
                    <thead>
                        <tr>
                            <th>Algoritma</th>
                            <th>Perbandingan Terukur</th>
                            <th>Perkiraan Teoretis</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr><td>Bubble Sort</td><td>498.834</td><td>Mendekati n(n−1)/2 = 499.500</td></tr>
                        <tr><td>Selection Sort</td><td>499.500</td><td>Tepat n(n−1)/2 = 499.500</td></tr>
                        <tr><td>Insertion Sort</td><td><span class="mono-bold">256.316</span></td><td>Sekitar n²/4 = 250.000</td></tr>
                    </tbody>
                </table>
            </div>
            <p class="section-desc" style="max-width: 100%; text-align: left;">Meskipun jumlah perbandingan Bubble Sort (498.834) hampir sama dengan Selection Sort (499.500), Bubble Sort jauh lebih lambat karena melakukan ~250.000 kali operasi pertukaran elemen (inversi) pada data acak, sedangkan Selection Sort maksimal hanya 999 pertukaran.</p>
        </div>
    </section>

    <!-- BAND 4: FITUR INTERAKTIF "COBA SENDIRI" (LIGHT) -->
    <section id="coba-demo" class="band-light">
        <div class="container">
            <div class="section-header">
                <div class="section-kicker">Interactive Live Simulator</div>
                <div class="section-title-wrap">
                    <h2>Coba Sendiri</h2>
                </div>
                <p class="section-desc">Masukkan angka atau buat data acak, lalu bandingkan ketiga algoritma langsung di browser. Waktu bergantung pada perangkatmu, jadi yang penting adalah perbandingan antar-algoritmanya.</p>
            </div>

            <div class="demo-box">
                <!-- Input Angka Manual -->
                <div class="demo-group">
                    <label class="demo-label" for="customNumbers">Angka sendiri (pisahkan dengan spasi atau koma)</label>
                    <textarea id="customNumbers" class="demo-textarea" placeholder="Contoh: 45, 12, 89, 3, 27">1, 2, 3, 4, 5</textarea>
                </div>
                <div class="demo-group">
                    <button type="button" class="btn-dark-solid" onclick="compareCustom()">Bandingkan angka saya</button>
                </div>

                <!-- Generator Data Acak -->
                <div class="demo-row">
                    <div>
                        <label class="demo-label" for="randomSize">Atau buat data acak, ukuran (n, maks 20.000)</label>
                        <input type="number" id="randomSize" class="demo-input" value="1000" min="2" max="20000">
                    </div>
                    <div>
                        <label class="demo-label" for="dataCondition">Kondisi data</label>
                        <select id="dataCondition" class="demo-select">
                            <option value="random" selected>Acak</option>
                            <option value="sorted">Terurut</option>
                            <option value="reversed">Terbalik</option>
                        </select>
                    </div>
                    <div>
                        <button type="button" class="btn-red" onclick="compareGenerated()">Buat &amp; bandingkan</button>
                    </div>
                </div>

                <!-- Status & Bar Hasil -->
                <div id="demoStatus" class="demo-status">1.000 angka berhasil diurutkan.</div>

                <div class="demo-results">
                    <div class="algo-row">
                        <div class="algo-name bubble">Bubble Sort</div>
                        <div>
                            <div class="bar-track"><div id="bar-bubble" class="bar-fill bubble" style="width: 100%;"></div></div>
                            <div id="meta-bubble" class="algo-meta">-</div>
                        </div>
                    </div>

                    <div class="algo-row">
                        <div class="algo-name selection">Selection Sort</div>
                        <div>
                            <div class="bar-track"><div id="bar-selection" class="bar-fill selection" style="width: 70%;"></div></div>
                            <div id="meta-selection" class="algo-meta">-</div>
                        </div>
                    </div>

                    <div class="algo-row">
                        <div class="algo-name insertion">Insertion Sort</div>
                        <div>
                            <div class="bar-track"><div id="bar-insertion" class="bar-fill insertion" style="width: 35%;"></div></div>
                            <div id="meta-insertion" class="algo-meta">-</div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </section>

    <!-- BAND 5: KESIMPULAN (DARK) -->
    <section id="kesimpulan" class="band-dark">
        <div class="container">
            <div class="section-header">
                <div class="section-kicker">Final Synthesis</div>
                <div class="section-title-wrap">
                    <h2>Kesimpulan Pengujian</h2>
                </div>
            </div>

            <div class="conclusion-grid">
                <div class="editorial-card">
                    <div class="card-kicker">Poin 01</div>
                    <h3>Pertumbuhan Kuadratik O(n²)</h3>
                    <div class="red-bar"></div>
                    <p>Pada data acak, waktu eksekusi ketiga algoritma tumbuh secara kuadratik terhadap ukuran data. Bubble Sort paling lambat, sekitar 2,3 sampai 2,5 kali lebih lama dibanding Selection Sort dan Insertion Sort.</p>
                </div>
                <div class="editorial-card">
                    <div class="card-kicker">Poin 02</div>
                    <h3>Sensitivitas Data Terurut</h3>
                    <div class="red-bar"></div>
                    <p>Kondisi data sangat memengaruhi kecepatan. Pada data terurut, Bubble Sort (dengan penanda pertukaran) dan Insertion Sort jauh lebih cepat karena berperilaku O(n), sedangkan Selection Sort tidak terpengaruh.</p>
                </div>
                <div class="editorial-card">
                    <div class="card-kicker">Poin 03</div>
                    <h3>Keunggulan di Data Terbalik</h3>
                    <div class="red-bar"></div>
                    <p>Pada data terurut terbalik (worst case), Selection Sort menjadi yang tercepat karena jumlah pertukarannya paling sedikit dibandingkan pergeseran dan pertukaran pada dua algoritma lainnya.</p>
                </div>
                <div class="editorial-card">
                    <div class="card-kicker">Poin 04</div>
                    <h3>Validasi Teori Empiris</h3>
                    <div class="red-bar"></div>
                    <p>Hasil pengukuran waktu dan jumlah perbandingan antar-elemen konsisten dengan analisis teoretis, membuktikan bahwa pengukuran empiris dapat dipakai untuk memvalidasi teori kompleksitas.</p>
                </div>
            </div>
        </div>
    </section>

    <footer>
        Kelompok 26 — Pellinore • Tugas Besar Fortran 2026
    </footer>

    <script>
        Chart.defaults.font.family = "'Plus Jakarta Sans', sans-serif";

        // Grafik 1
        new Chart(document.getElementById('randomChart'), {
            type: 'line',
            data: {
                labels: ['100', '500', '1.000', '2.500', '5.000'],
                datasets: [
                    {
                        label: 'Bubble Sort',
                        data: [0.31, 6.39, 27.43, 193.64, 825.30],
                        borderColor: '#36444f',
                        backgroundColor: 'rgba(54, 68, 79, 0.14)',
                        fill: true,
                        tension: 0.25,
                        borderWidth: 3
                    },
                    {
                        label: 'Selection Sort',
                        data: [0.11, 3.04, 12.18, 83.29, 327.49],
                        borderColor: '#18181b',
                        backgroundColor: '#18181b',
                        borderDash: [6, 4],
                        tension: 0.25,
                        borderWidth: 2.5
                    },
                    {
                        label: 'Insertion Sort',
                        data: [0.10, 2.93, 11.80, 80.66, 352.91],
                        borderColor: '#9ca3af',
                        backgroundColor: '#9ca3af',
                        borderDash: [2, 3],
                        tension: 0.25,
                        borderWidth: 2.5
                    }
                ]
            },
            options: {
                responsive: true,
                plugins: {
                    title: { display: true, text: 'GAMBAR 1. WAKTU EKSEKUSI PADA DATA ACAK (MS)', color: '#18181b', font: { family: 'Montserrat', size: 13, weight: 'bold' } },
                    tooltip: { mode: 'index', intersect: false }
                },
                scales: {
                    x: { title: { display: true, text: 'Ukuran Data (n)', color: '#52525b' }, grid: { color: '#f4f4f5' } },
                    y: { title: { display: true, text: 'Waktu Eksekusi (ms)', color: '#52525b' }, grid: { color: '#f4f4f5' }, beginAtZero: true }
                }
            }
        });

        // Grafik 2
        new Chart(document.getElementById('casesChart'), {
            type: 'bar',
            data: {
                labels: ['Terurut (Best Case)', 'Acak (Average Case)', 'Terbalik (Worst Case)'],
                datasets: [
                    {
                        label: 'Bubble Sort',
                        data: [0.07, 113.45, 157.24],
                        backgroundColor: '#36444f',
                        borderColor: '#647d91',
                        borderWidth: 1.5,
                        borderRadius: 2
                    },
                    { label: 'Selection Sort', data: [48.38, 46.81, 51.87], backgroundColor: '#ffffff', borderRadius: 2 },
                    { label: 'Insertion Sort', data: [0.12, 58.49, 89.16], backgroundColor: '#71717a', borderRadius: 2 }
                ]
            },
            options: {
                responsive: true,
                plugins: {
                    title: { display: true, text: 'GAMBAR 2. PERBANDINGAN WAKTU PADA BERBAGAI KONDISI DATA (N = 2.000)', color: '#ffffff', font: { family: 'Montserrat', size: 13, weight: 'bold' } },
                    legend: { labels: { color: '#e4e4e7' } }
                },
                scales: {
                    x: { ticks: { color: '#a1a1aa' }, grid: { color: '#32363a' } },
                    y: { ticks: { color: '#a1a1aa' }, title: { display: true, text: 'Waktu Eksekusi (ms)', color: '#a1a1aa' }, grid: { color: '#32363a' }, beginAtZero: true }
                }
            }
        });

        // === LOGIKA SIMULATOR "COBA SENDIRI" ===
        function runBubble(arr) {
            const a = arr.slice();
            let comp = 0, swaps = 0;
            const n = a.length;
            const t0 = performance.now();
            for (let i = 0; i < n - 1; i++) {
                let swapped = false;
                for (let j = 0; j < n - 1 - i; j++) {
                    comp++;
                    if (a[j] > a[j + 1]) {
                        const tmp = a[j]; a[j] = a[j + 1]; a[j + 1] = tmp;
                        swaps++;
                        swapped = true;
                    }
                }
                if (!swapped) break;
            }
            const t1 = performance.now();
            return { time: t1 - t0, comp, swaps };
        }

        function runSelection(arr) {
            const a = arr.slice();
            let comp = 0, swaps = 0;
            const n = a.length;
            const t0 = performance.now();
            for (let i = 0; i < n - 1; i++) {
                let m = i;
                for (let j = i + 1; j < n; j++) {
                    comp++;
                    if (a[j] < a[m]) m = j;
                }
                if (m !== i) {
                    const tmp = a[i]; a[i] = a[m]; a[m] = tmp;
                    swaps++;
                }
            }
            const t1 = performance.now();
            return { time: t1 - t0, comp, swaps };
        }

        function runInsertion(arr) {
            const a = arr.slice();
            let comp = 0, shifts = 0;
            const n = a.length;
            const t0 = performance.now();
            for (let i = 1; i < n; i++) {
                const key = a[i];
                let j = i - 1;
                while (j >= 0) {
                    comp++;
                    if (a[j] > key) {
                        a[j + 1] = a[j];
                        shifts++;
                        j--;
                    } else {
                        break;
                    }
                }
                a[j + 1] = key;
            }
            const t1 = performance.now();
            return { time: t1 - t0, comp, swaps: shifts };
        }

        function formatMs(ms) {
            if (ms < 0.05) return ms.toFixed(3).replace('.', ',') + ' ms';
            return Number(ms.toFixed(1)).toLocaleString('id-ID') + ' ms';
        }

        function executeBenchmark(data) {
            const res = {
                bubble: runBubble(data),
                selection: runSelection(data),
                insertion: runInsertion(data)
            };

            document.getElementById('demoStatus').textContent =
                `${data.length.toLocaleString('id-ID')} angka berhasil diurutkan.`;

            // Tentukan pemenang tercepat
            const times = [
                { key: 'bubble', val: res.bubble.time },
                { key: 'selection', val: res.selection.time },
                { key: 'insertion', val: res.insertion.time }
            ];
            times.sort((a, b) => a.val - b.val);
            const fastestKey = times[0].key;

            // Skala lebar bar
            const maxTime = Math.max(res.bubble.time, res.selection.time, res.insertion.time);
            const maxComp = Math.max(res.bubble.comp, res.selection.comp, res.insertion.comp, 1);

            ['bubble', 'selection', 'insertion'].forEach(key => {
                const item = res[key];
                let pct = maxTime > 0.05
                    ? (item.time / maxTime) * 100
                    : (item.comp / maxComp) * 100;
                pct = Math.max(pct, 6);

                document.getElementById(`bar-${key}`).style.width = `${pct}%`;

                const timeText = key === fastestKey
                    ? `<span class="fastest-badge">${formatMs(item.time)} • tercepat</span>`
                    : `${formatMs(item.time)}`;

                document.getElementById(`meta-${key}`).innerHTML =
                    `${timeText} · ${item.comp.toLocaleString('id-ID')} perbandingan · ${item.swaps.toLocaleString('id-ID')} pertukaran/pergeseran`;
            });
        }

        function compareCustom() {
            const raw = document.getElementById('customNumbers').value.trim();
            if (!raw) {
                alert('Masukkan minimal 2 angka terlebih dahulu.');
                return;
            }
            const nums = raw.split(/[\s,]+/).map(Number).filter(n => !isNaN(n));
            if (nums.length < 2) {
                alert('Masukkan minimal 2 angka valid yang dipisahkan spasi atau koma.');
                return;
            }
            executeBenchmark(nums);
        }

        function compareGenerated() {
            let n = parseInt(document.getElementById('randomSize').value, 10);
            if (isNaN(n) || n < 2) n = 100;
            if (n > 20000) {
                n = 20000;
                document.getElementById('randomSize').value = 20000;
            }
            const cond = document.getElementById('dataCondition').value;
            let data = Array.from({ length: n }, () => Math.floor(Math.random() * 100000));

            if (cond === 'sorted') {
                data.sort((a, b) => a - b);
            } else if (cond === 'reversed') {
                data.sort((a, b) => b - a);
            }
            executeBenchmark(data);
        }

        // Jalankan simulasi awal otomatis saat halaman dimuat
        window.addEventListener('DOMContentLoaded', compareGenerated);
    </script>
</body>
</html>