package erratum

// CARA : note that i can set the variable for the return at the top (var err)
func Use(o ResourceOpener, input string) (err error) {
	var res Resource

	for res, err = o(); err != nil; res, err = o() {
		if _, ok := err.(TransientError); !ok {
			return err
		}
	}

	defer res.Close()
	defer func() {
		if rec := recover(); rec != nil {
			switch e := rec.(type) {
			case FrobError:
				res.Defrob(e.defrobTag)
				err = e
			case error:
				err = e
			}
		}
	}()
	res.Frob(input)

	return err
}