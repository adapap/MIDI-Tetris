import mido
from pynput.keyboard import Key, Controller
kb = Controller()
keys = {
    48: 'z',
    50: 'x',
    52: 'c',
    60: Key.space,
    62: Key.down,
    64: Key.left,
    67: Key.right,
}
with mido.open_input() as port:
    for message in port:
        if message.type == 'note_on':
            key = keys.get(message.note)
            if key is not None:
                if message.velocity > 0:
                    kb.press(key)
                else:
                    kb.release(key)