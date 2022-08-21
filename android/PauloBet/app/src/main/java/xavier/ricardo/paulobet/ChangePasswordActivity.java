package xavier.ricardo.paulobet;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Intent;
import android.os.Bundle;
import android.view.View;
import android.widget.EditText;
import android.widget.Toast;

import java.io.IOException;

import okhttp3.Response;
import xavier.ricardo.paulobet.tasks.ChangePasswordTask;

public class ChangePasswordActivity extends AppCompatActivity {

    private String user;
    private String token;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_change_password);
        Intent intent = getIntent();
        user = intent.getStringExtra("user");
        token = intent.getStringExtra("token");
    }

    public void confirm(View v) {
        EditText tpNewPassword = findViewById(R.id.tpNewPassword);
        EditText tpConfirmPassword = findViewById(R.id.tpConfirmPassword);
        if (!tpNewPassword.getText().toString().equals(tpConfirmPassword.getText().toString())) {
            Toast.makeText(this, "Senhas diferentes", Toast.LENGTH_LONG).show();
            return;
        }
        String password = tpNewPassword.getText().toString().trim();
        new ChangePasswordTask(this, user, password, token).execute();
    }

    public void onTaskResponse(Response response) throws IOException {
        if (response == null) {
            Toast.makeText(this, "Sem resposta so servidor. Tente novamente mais tarde.", Toast.LENGTH_LONG).show();
            return;
        }
        if (response.code() == 204) {
            Toast.makeText(this, "Senha alterada com sucesso", Toast.LENGTH_LONG).show();
            finish();
        } else {
            Toast.makeText(this, "ERRO status " + response.code(), Toast.LENGTH_LONG).show();
        }
    }
}