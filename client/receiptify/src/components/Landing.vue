<template>
  <div id="landing">
    <div class="create_account">
      <v-col cols="auto">
        <v-dialog transition="dialog-bottom-transition" max-width="600">
          <template v-slot:activator="{ on, attrs }">
            <v-chip
              style="width:120px; height:40px; text-align:center;"
              color="blue"
              v-bind="attrs"
              v-on="on"
              outlined
              >Create Account >
            </v-chip>
          </template>
          <template v-slot:default="dialog">
            <v-card>
              <v-toolbar class="account_registration_header" dark
                >Lets get an account setup.</v-toolbar
              >
              <v-card-text>
                <v-form ref="form" v-model="valid" lazy-validation>
                  <v-text-field
                    v-model="name"
                    :counter="10"
                    :rules="nameRules"
                    label="Name"
                    required
                  ></v-text-field>

                  <v-text-field
                    v-model="email"
                    :rules="emailRules"
                    label="E-mail"
                    required
                  ></v-text-field>

                  <v-select
                    v-model="select"
                    :items="items"
                    :rules="[(v) => !!v || 'Item is required']"
                    label="Item"
                    required
                  ></v-select>

                  <v-checkbox
                    v-model="checkbox"
                    :rules="[(v) => !!v || 'You must agree to continue!']"
                    label="Do you agree?"
                    required
                  ></v-checkbox>

                  <v-btn
                    :disabled="!valid"
                    color="success"
                    class="mr-4"
                    @click="validate"
                  >
                    Validate
                  </v-btn>

                  <v-btn color="error" class="mr-4" @click="reset">
                    Reset Form
                  </v-btn>

                  <v-btn color="warning" @click="resetValidation">
                    Reset Validation
                  </v-btn>
                </v-form>
              </v-card-text>
              <v-card-actions class="justify-end">
                <v-btn text @click="dialog.value = false">Close</v-btn>
              </v-card-actions>
            </v-card>
          </template>
        </v-dialog>
      </v-col>
    </div>
    <div class="digitize_label">
      <img class="d_label" src="@/assets/digitizeLabel.png" />
    </div>
    <div class="digitize">
      <router-link to="/digitize"
        ><img src="@/assets/digitize.png" />
      </router-link>
    </div>

    <div class="content">
      <event-hub></event-hub>
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
export default {
  name: "landing",
  data: () => ({
    valid: true,
    name: "",
    nameRules: [
      (v) => !!v || "Name is required",
      (v) => (v && v.length <= 10) || "Name must be less than 10 characters",
    ],
    email: "",
    emailRules: [
      (v) => !!v || "E-mail is required",
      (v) => /.+@.+\..+/.test(v) || "E-mail must be valid",
    ],
    select: null,
    items: ["Item 1", "Item 2", "Item 3", "Item 4"],
    checkbox: false,
  }),
  methods: {
    validate() {
      this.$refs.form.validate();
    },
    reset() {
      this.$refs.form.reset();
    },
    resetValidation() {
      this.$refs.form.resetValidation();
    },
  },
};
</script>

<style>
#landing {
  position: relative;
  width: 100%;
  height: 100vh;
}

.account_registration_header {
  background-color: rgba(47, 44, 54, 0.397);
}

@media (max-width: 599px) {
  .create_account {
    position: absolute;
    left: 60%;
    top: 50%;
    bottom: 50%;
    width: 140px;
    height: 50px;
  }
  .digitize {
    position: absolute;
    object-fit: contain;
    width: 100%;
    left: 55%;
    bottom: 0%;
  }
  .digitize_label {
    position: absolute;
    object-fit: contain;
    width: auto;
    left: 15%;
    bottom: 7%;
  }

  .d_label {
    width: 200px !important;
  }
}
@media (min-width: 600px) {
  .create_account {
    position: absolute;
    left: 60%;
    top: 50%;
    bottom: 50%;
    width: 140px;
    height: 50px;
  }

  .digitize {
    position: absolute;
    object-fit: contain;
    width: 100%;
    left: 60%;
    bottom: 2%;
  }
  .digitize_label {
    position: absolute;
    object-fit: contain;
    width: auto;
    left: 15%;
    bottom: 7%;
  }

  .d_label {
    width: 200px !important;
  }
}

@media (min-width: 1100px) {
  .create_account {
    position: absolute;
    left: 80%;
    top: 50%;
    bottom: 50%;
    width: 140px;
    height: 50px;
  }
  .digitize {
    position: absolute;
    object-fit: contain;
    width: 100%;
    left: 50%;
    bottom: 2%;
  }
  .digitize_label {
    position: absolute;
    object-fit: contain;
    width: auto;
    left: 15%;
    bottom: 7%;
  }

  .d_label {
    width: 400px !important;
  }
}
</style>
