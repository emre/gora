# gora  [![Clone in Koding](http://kbutton.org/clone.png?v=2)](http://kbutton.org/emre/gora)

command line seslisozluk client for *tr-en*, *en-tr* translations.

### installation
```bash
$ go get github.com/emre/gora
```
### usage
```bash
$ gora [word_to_translate]
```

**english to turkish**
    
```bash
$ gora peace
1. barış. 2. sessiz olun. 3. iç huzuru. 4. huzur. 5. hazar.
```

**turkish to english**
```bash
$ gora barış en_tr
1. peace. 2. concord. 3. 1.peace. 4. peacetime. 5. reconciliation.
```

**turkish to turkish**

```bash
$ gora emre tr_tr
1. Aşık, mübtela, vurgun. 2. Ak gözlü, beyaz gözlü. 3. Ak gözlü, beyaz gözl. 
```
