Show why it´s sometimes neccessary to use mutexes (or something similiar, not going into detail about other options like CRDTs, lock-free algorithms, etc.).

If you run this program (hint: it has paramaters, see "-h"), the correct result should be "100". However, due to races the default run produces incorrect results.
