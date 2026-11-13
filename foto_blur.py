import cv2
import mediapipe as mp

mp_hands = mp.solutions.hands
hands = mp_hands.Hands(
    static_image_mode=False,
    max_num_hands=1,
    min_detection_confidence=0.5,
    min_tracking_confidence=0.5
)
def finger_up(tip, pip, landmarks):
    return landmarks[tip].y < landmarks[pip].y

def is_peace(landmarks):
    telunjuk = finger_up(8, 6, landmarks)
    tengah = finger_up(12, 10, landmarks)
    manis = finger_up(16, 14, landmarks)
    kelingking = finger_up(20, 18, landmarks);
    
    return telunjuk and tengah and not manis and not kelingking

cap = cv2.VideoCapture(0)

print("Kamera menyala! Kasih pose ✌️ ke arah kamera buat nge-blur.")
print("Tekan tombol 'q' pada keyboard untuk keluar.")

while True:
    sukses, frame = cap.read()
    if not sukses:
        break

    frame = cv2.flip(frame, 1)
    frame_rgb = cv2.cvtColor(frame, cv2.COLOR_BGR2RGB)
    
    hasil_ai = hands.process(frame_rgb)

    if hasil_ai.multi_hand_landmarks:
        for tangan in hasil_ai.multi_hand_landmarks:
            if is_peace(tangan.landmark):
                frame = cv2.GaussianBlur(frame, (51, 51), 0)

    cv2.imshow("Kamera Pintar AI", frame)

    if cv2.waitKey(1) & 0xFF == ord('q'):
        break

cap.release()
cv2.destroyAllWindows()